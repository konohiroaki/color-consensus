import React, {Component} from "react";
import {CopyToClipboard} from "react-copy-to-clipboard";

class Header extends Component {

    render() {
        console.log("rendering header");
        const userId = this.props.userId;
        const button = userId === undefined || userId === null
                       ? Header.signUpLoginButton()
                       : Header.userIdButton(userId);
        return (
            <nav className="navbar navbar-dark bg-dark border-bottom border-secondary" style={this.props.style}>
                <a className="navbar-brand" href="/">Color Consensus</a>
                {button}
            </nav>
        );
    }

    static signUpLoginButton() {
        // TODO: get modal from props to depart from tight couple?
        return (
            <button className="btn btn-outline-light"
                    data-toggle="modal" data-target="#signup-login-modal">
                Sign Up / Login
            </button>
        );
    }

    static userIdButton(userId) {
        // TODO: show a dialog(?) to notify a user that id is copied to their clipboard.
        return (
            <CopyToClipboard text={userId}>
                <button className="btn btn-outline-secondary">
                    ID: {userId}
                </button>
            </CopyToClipboard>
        );
    }
}

export default Header;