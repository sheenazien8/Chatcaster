import { useEffect, useState } from 'react';
import './App.css';
import { Config } from '../wailsjs/go/main/App';
import { EventsOn, EventsOff } from '../wailsjs/runtime/runtime';

function App() {
    const [text, setText] = useState('');
    const [prevText, setPrevText] = useState('');
    const [config, setConfig] = useState({});

    const handleText = (char) => {
        if (char == 'Backspace') {
            setText((prev) => prev.slice(0, -1));
            return;
        }
        if (char === 'Shift' || char === 'Control' || char === 'Alt' || char === 'Meta' || char === 'Tab' || char == 'Esc') {
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
        const loadConfig = async () => {
            setConfig(await Config());
        };

        loadConfig();
    }, []);


    useEffect(() => {
        const handleKeyDown = (event) => {
            handleText(event.key);
        };

        console.log('config', config);
        if (config.BubbleChat) {
            document.documentElement.style.setProperty('--bubble-color', config.BubbleChat.Style.BackgroundColor);
            document.documentElement.style.setProperty('--font-color', config.BubbleChat.Style.FontColor);
            document.documentElement.style.setProperty('--font-size', config.BubbleChat.Style.FontSize);
            if (config.BubbleChat.ShowKey.Mode == "focus") {
                window.addEventListener('keydown', handleKeyDown);

                return () => {
                    window.removeEventListener('keydown', handleKeyDown);
                };
            } else {
                EventsOn('keyPressed', (data) => {
                    handleText(data);
                });
                return () => {
                    EventsOff('keyPressed');
                };
            }
        }
    }, [config]);


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

