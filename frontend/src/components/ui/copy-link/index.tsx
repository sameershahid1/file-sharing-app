import { FC, useState } from 'react';

type Props = {
    link: string;
}


const CopyLink: FC<Props> = ({ link }) => {
    const [copied, setCopied] = useState(false);

    const copyToClipboard = () => {
        navigator.clipboard.writeText(link)
            .then(() => {
                setCopied(true);
            })
    };

    return (
        <div className="flex items-center space-x-2 p-4 border rounded shadow-md">
            <input
                className="flex-grow p-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                type="text"
                value={link}
                readOnly
            />
            <button
                onClick={copyToClipboard}
                className={`px-4 py-2 rounded transition-colors ${copied
                    ? 'bg-green-500 text-white hover:bg-green-600'
                    : 'bg-blue-500 text-white hover:bg-blue-600'
                    }`}
            >
                {copied ? 'Copied!' : 'Copy'}
            </button>
        </div>
    );
};



export default CopyLink
