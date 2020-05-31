import React from 'react';
import ReactDOM from 'react-dom';
// Basic React Components.

type MyButtonProps = {
    name:string;
}

type MyButtonState = {
    pressedCount:number;
}

class MyButton extends React.Component<MyButtonProps, MyButtonState> {
    constructor(props: MyButtonProps){
        super(props);
        this.state = {pressedCount:0};
    }

    onClickAction: () => void = () => {
        alert(`clicked! ${this.state.pressedCount} times.`);
        this.setState({pressedCount: this.state.pressedCount +1 });
    }

    render = () => {
        return (
            <div>
                <button onClick={this.onClickAction}>{this.props.name}</button>
            </div>
        )
    }
}

ReactDOM.render(<MyButton name={`Count Up`}></MyButton>, document.getElementById('original-button'));
