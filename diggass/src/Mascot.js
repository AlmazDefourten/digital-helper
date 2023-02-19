import React from 'react';
import './Mascot.css';

class Mascot extends React.Component{
    constructor() {
        super()
        this.state = {
            name: ''
        }
    }
    handleChange = (event) => {
        this.setState({[event.target.name]: event.target.value});
      }
     
      handleSubmit = (event) => {
        alert('A form was submitted: ' + this.state);
     
        fetch('http://localhost:8000/api/v1/chat/send?' + new URLSearchParams({
            question: document.getElementById('text').value
        }), {
            method: 'GET',
          }).then(response => response.json()).then(function(response) {
            console.log(response)
            const messages = [response.data.answer];
            const messageList = document.getElementById('message-list');
            messages.forEach(message => {
                const div = document.createElement('div');
                div.style.border = '1px solid lightgreen';
                div.style.backgroundColor = 'lightgreen';
                div.style.padding = '10px';
                div.textContent = response.data.answer;
                messageList.appendChild(div);
            });
          });

        event.preventDefault();
    }
    render() {
        return(
            <div className='mas'>
                <div className='mascot'></div>
                <div className='chat'>
                    <div className='okno'>
                        <form onSubmit={this.handleSubmit}>
                            <Out messages={this.state.messages}></Out>
                            <div className='chat_form'>
                                <input id='text' className='text_in' placeholder='Поле ввода' type="text" value={this.state.value} name="name" onChange={this.handleChange}></input>
                                <button className='but_chat' value="Submit"></button>
                            </div>
                        </form>
                    </div>
                    <button className='clear'>X</button>
                </div>
            </div>
        );
    }
}
class Out extends React.Component {
    render() {
        return (
            <ul className="message-list" id="message-list">

            </ul>
        )
    }
}

export default Mascot;