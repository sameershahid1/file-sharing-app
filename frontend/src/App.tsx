
//Packages
import { useEffect, useState } from 'react';

//Go Bind methods
import { GetMultiAddress } from "../wailsjs/go/app/App";
import CopyLink from './components/ui/copy-link';


function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    useEffect(() => {
        GetMultiAddress().then(updateResultText)
    }, [])

    function greet() {
    }

    return (
        <div className='w-80 mx-auto'>
            <p className="text-red-800 bg-white">{resultText}</p>
            <div className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text" />
                <button className="btn" onClick={greet}>Greet</button>
            </div>
            <CopyLink
                link={resultText}
            />
        </div>
    )
}

export default App
