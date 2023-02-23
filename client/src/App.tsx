import {MessageInput} from './components/MessageInput'
import {MessageList} from './components/MessageList'

export const App = () => {
    return (
        <div>
            <h1>Real Time Chat</h1>
            <MessageInput/>
            <MessageList/>
        </div>
    )
}