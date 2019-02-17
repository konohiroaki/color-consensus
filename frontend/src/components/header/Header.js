import React, {Component} from "react";
import axios from "axios";
import $ from "jquery";

class Header extends Component {

    constructor(props) {
        super(props);
        this.state = {};
    }

    render() {
        const userId = this.state.userId;
        const button = userId === undefined || userId === null
                       ? null
            // TODO: click button to copy userId into clipboard
                       : <button className="btn btn-outline-secondary">ID: {userId}</button>;
        const modal = userId === null
                      ? this.modal()
                      : null;
        console.log("rendering header");
        return (
            <nav className="navbar navbar-dark bg-dark border-bottom border-secondary" style={this.props.style}>
                <a className="navbar-brand" href="/">Color Consensus</a>
                {button}
                {modal}
            </nav>
        );
    }

    modal() {
        return (
            <div className="modal fade" tabIndex="-1" id="signup-login-modal">
                <div className="modal-dialog modal-dialog-centered">
                    <div className="modal-content bg-dark">
                        <div className="modal-body">
                            <ul className="nav nav-tabs" role="tablist"
                                style={{
                                    // deny .modal-body's 1em padding for right and left.
                                    marginLeft: "-1em", marginRight: "-1em",
                                    // deny .nav style and use <ul> style
                                    marginBottom: "1rem", paddingLeft: "40px"
                                }}>
                                <li className="nav-item">
                                    <a className="nav-link active" href="#signup-tab" data-toggle="tab" role="tab">Sign Up</a>
                                </li>
                                <li className="nav-item">
                                    <a className="nav-link" href="#login-tab" data-toggle="tab" role="tab">Login</a>
                                </li>
                            </ul>
                            <div className="tab-content" id="myTabContent">
                                {/* TODO: implement tab contents*/}
                                <div className="tab-pane fade show active" id="signup-tab" role="tabpanel">sign up tab</div>
                                <div className="tab-pane fade" id="login-tab" role="tabpanel">login tab</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }

    componentDidMount() {
        axios.get("http://localhost:5000/api/v1/users/presence")
            .then(({data}) => {
                // TODO: show user id on header
                console.log("user present", data);
                this.setState({userId: data.userID});
            })
            .catch(() => {
                // TODO: show modal to login or sign up.
                console.log("user not present");
                this.setState({userId: null});
            });
    }

    // TODO: even after closing the modal, give a way or force user to login/signup before doing anything else.
    componentDidUpdate() {
        if (this.state.userId === null) {
            $("#signup-login-modal").modal();
        }
    }
}

export default Header;