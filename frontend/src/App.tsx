
//Packages
import { FormEvent, useEffect, useState } from 'react';

//Lib
import CopyLink from './components/ui/copy-link';

//Go Bind methods
import { FileShare, GetMultiAddress, OpenFileDialogBox } from "../wailsjs/go/app/App";


function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const updateResultText = (result: string) => setResultText(result);
    const [targetHostLink, setTargetHostLink] = useState<string>("")

    useEffect(() => {
        GetMultiAddress().then(updateResultText)
    }, [])


    const handleSelect = async () => {
        const filePath = await OpenFileDialogBox()
        console.log(filePath)
        if (targetHostLink.length !== 0) {
            const response = await FileShare(targetHostLink, filePath)
            console.log(response)
        }
    }

    const handleConnectTargetHost = (evt: FormEvent) => {
        evt.preventDefault()
        const formElement = evt.currentTarget as HTMLFormElement
        const hostField = formElement.elements[0] as HTMLInputElement
        setTargetHostLink(hostField.value)
    }


    return (
        <div className='bg-gray-700 h-screen w-screen pt-10'>
            <div className='flex flex-col p-8  shadow-md rounded-lg border border-gray-400 gap-2 w-[84vw] mx-auto text-white'>
                <form onSubmit={handleConnectTargetHost} >
                    <label className="block mb-1 text-sm font-semibold">
                        Target Peer Address
                    </label>
                    <input
                        className="w-[40vw] bg-transparent text-white placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
                        placeholder="Target peer link"
                        required={true}
                        type={"text"}
                    />
                    <button
                        className="rounded-md cursor-pointer bg-slate-800 py-2 px-4 border border-transparent text-center text-sm text-white transition-all shadow-md hover:shadow-lg focus:bg-slate-700 focus:shadow-none active:bg-slate-700 hover:bg-slate-700 active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none ml-2"
                        type="submit"
                    >
                        Connect
                    </button>
                </form>

                <div className='flex items-center gap-4 justify-between'>
                    <CopyLink
                        link={resultText}
                    />
                    <button
                        className="flex-1 text-gray-500 font-medium text-lg bg-gray-100 file:cursor-pointer cursor-pointer file:border-0 file:py-3 file:px-4 file:mr-4 file:bg-gray-800 file:hover:bg-gray-700 file:text-white rounded"
                        onClick={handleSelect}
                        type="button"
                    >
                        Upload
                    </button>
                </div>
                <div className='flex items-center gap-10 '>
                    <p className='font-bold text-[1rem]'>Loading: </p>
                    <div
                        className="flex-start flex h-4 w-full overflow-hidden rounded-full bg-gray-50 font-sans text-xs font-medium">
                        <div
                            className="flex items-center justify-center h-full overflow-hidden text-white break-all bg-red-500 rounded-full"
                            style={{ width: "50%" }}
                        >
                            25% Completed
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default App
