import React, {Component} from "react";
import {CopyToClipboard} from "react-copy-to-clipboard";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faGithub} from "@fortawesome/free-brands-svg-icons";
import {connect} from "react-redux";

class Header extends Component {

    constructor(props) {
        super(props);
    }

    render() {
        console.log("rendering header", this.props.userId);

        return <nav className="navbar navbar-dark bg-dark border-bottom border-secondary" style={this.props.style}>
            <HeaderLeftPart/>
            <HeaderRightPart userId={this.props.userId} loginModalRef={this.props.loginModalRef}/>
        </nav>;
    }
}

const HeaderLeftPart = () => (
    <div>
        <a className="navbar-brand" href="/">Color Consensus</a>
        <GithubLink/>
    </div>
);

const GithubLink = () => (
    // FIXME: placement of icon is not good. place it a bit lower.
    <a href="https://github.com/konohiroaki/color-consensus" className="text-light">
        <FontAwesomeIcon icon={faGithub} size="2x"/>
    </a>
);

const HeaderRightPart = ({userId, loginModalRef}) => {
    return userId === undefined || userId === null
           ? <SignUpLoginModalButton loginModalRef={loginModalRef}/>
           : <UserIdCopyButton userId={userId}/>;
};

const SignUpLoginModalButton = ({loginModalRef}) => (
    <button className="btn btn-outline-light" onClick={() => loginModalRef.current.openLoginModal()}>
        Sign Up / Login
    </button>
);

// TODO: show a dialog(?) to notify a user that id is copied to their clipboard.
const UserIdCopyButton = ({userId}) => (
    <CopyToClipboard text={userId}>
        <button className="btn btn-outline-secondary">
            ID: {userId}
        </button>
    </CopyToClipboard>
);

const mapStateToProps = state => ({
    userId: state.user.id,
});

export default connect(mapStateToProps)(Header);
