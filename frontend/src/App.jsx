import { useEffect, useState } from 'react';
import './App.css';

function App() {
    const [text, setText] = useState('');
    const [prevText, setPrevText] = useState('');

    const handleText = (char) => {
        if (char == 'Backspace') {
            setText((prev) => prev.slice(0, -1));
            return;
        }

        setText((prev) => {
            if (char == 'Enter') {
                console.log('Enter', prev)
                setPrevText(prev);
                // setText('');
                return '';
            }
            return prev + char;
        });
    }

    useEffect(() => {
        const handleKeyDown = (event) => {
            if (event.key === 'Shift' || event.key === 'Control' || event.key === 'Alt' || event.key === 'Meta' || event.key === 'Tab') {
                return;
            }
            handleText(event.key);
        };

        window.addEventListener('keydown', handleKeyDown);

        return () => {
            window.removeEventListener('keydown', handleKeyDown);
        };
    }, []);


    useEffect(() => {
        if (text) {
            const timer = setTimeout(() => {
                setPrevText(text);
                setText('');
            }, 2000);

            return () => clearTimeout(timer);
        }
    }, [text]);

    useEffect(() => {
        if (prevText) {
            const timer = setTimeout(() => {
                setPrevText('');
            }, 4000);

            return () => clearTimeout(timer);
        }
    }, [prevText]);

    return (
        <div id="App">
            <div>
                {prevText != '' ? <p>{prevText}</p> : null}
            </div>
            <div>
                {text != '' ? <p>{text}</p> : null}
            </div>
        </div>
    )
}

export default App

