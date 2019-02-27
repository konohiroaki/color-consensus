import React, {Component} from "react";
import {CopyToClipboard} from "react-copy-to-clipboard";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faGithub} from "@fortawesome/free-brands-svg-icons";

class Header extends Component {

    constructor(props) {
        super(props);

        this.signUpLoginButton = this.signUpLoginButton.bind(this);
    }

    render() {
        console.log("rendering header");
        const userId = this.props.userId;
        const button = userId === undefined || userId === null
                       ? this.signUpLoginButton()
                       : Header.userIdButton(userId);
        return (
            <nav className="navbar navbar-dark bg-dark border-bottom border-secondary" style={this.props.style}>
                <div>
                    <a className="navbar-brand" href="/">Color Consensus</a>
                    <GithubLink/>
                </div>
                {button}
            </nav>
        );
    }

    signUpLoginButton() {
        return (
            <button className="btn btn-outline-light" onClick={() => this.props.loginModalRef.openLoginModal()}>
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

const GithubLink = () => (
    // FIXME: placement of icon is not good. place it a bit lower.
    <a href="https://github.com/konohiroaki/color-consensus" className="text-light">
        <FontAwesomeIcon icon={faGithub} size="2x"/>
    </a>
);

export default Header;
