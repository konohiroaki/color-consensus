import React, {Component} from "react";

class Header extends Component {

    render() {
        const userId = this.props.userId;
        const button = userId === undefined || userId === null
                       ? this.signUpLoginButton()
            // TODO: click button to copy userId into clipboard
                       : <button className="btn btn-outline-secondary">ID: {userId}</button>;
        console.log("rendering header");
        return (
            <nav className="navbar navbar-dark bg-dark border-bottom border-secondary" style={this.props.style}>
                <a className="navbar-brand" href="/">Color Consensus</a>
                {button}
            </nav>
        );
    }

    signUpLoginButton() {
        // TODO: get modal from props to depart from tight couple?
        return (<button className="btn btn-outline-light"
                        data-toggle="modal" data-target="#signup-login-modal">
            Sign Up / Login
        </button>);
    }
}

export default Header;