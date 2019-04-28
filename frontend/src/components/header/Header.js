import React, {Component} from "react";
import {CopyToClipboard} from "react-copy-to-clipboard";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faGithub} from "@fortawesome/free-brands-svg-icons";
import {connect} from "react-redux";
import ReactAwesomePopover from "react-awesome-popover";
import "react-awesome-popover/dest/react-awesome-popover.css";

class Header extends Component {

    constructor(props) {
        super(props);
    }

    render() {
        console.log("rendering [header]",
            "userId:", this.props.userId);

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
    <a href="https://github.com/konohiroaki/color-consensus" target="_blank" className="text-light">
        <FontAwesomeIcon icon={faGithub} size="lg"/>
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

const UserIdCopyButton = ({userId}) => (
    <ReactAwesomePopover action="hover" placement="left" arrow={false}>
        <CopyToClipboard text={userId}>
            <button className="btn btn-outline-light">
                ID: {userId}
            </button>
        </CopyToClipboard>
        <div>Click to copy</div>
    </ReactAwesomePopover>
);

const mapStateToProps = state => ({
    userId: state.user.id,
});

export default connect(mapStateToProps)(Header);
