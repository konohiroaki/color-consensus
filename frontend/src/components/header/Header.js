import React, {Component} from "react";
import axios from "axios";
import $ from "jquery";

class Header extends Component {

    constructor(props) {
        super(props);
        // TODO: lift up the login state because it's not header's state.
        this.state = {};

        this.handleSignUpClick = this.handleSignUpClick.bind(this);
        this.handleLoginClick = this.handleLoginClick.bind(this);
    }

    render() {
        const userId = this.state.userId;
        const button = userId === undefined || userId === null
                       ? this.signUpLoginButton()
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

    signUpLoginButton() {
        return (<button className="btn btn-outline-light"
                        data-toggle="modal" data-target="#signup-login-modal">
            Sign Up / Login
        </button>);
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
                                    {/* TODO: proper text color for nav link*/}
                                    <a className="nav-link active" href="#signup-tab" data-toggle="tab" role="tab">Sign Up</a>
                                </li>
                                <li className="nav-item">
                                    <a className="nav-link" href="#login-tab" data-toggle="tab" role="tab">Login</a>
                                </li>
                            </ul>
                            <div className="tab-content" id="myTabContent">
                                <div className="tab-pane fade show active" id="signup-tab" role="tabpanel">
                                    {/* TODO: get list from server and use select box */}
                                    Nationality:
                                    <input type="text" className="form-control" id="add-user-nationality" placeholder="Japan"
                                           onChange={e => this.setState({nationality: e.target.value})}/>
                                    {/* TODO: get list from server and use select box */}
                                    Gender:
                                    <input type="text" className="form-control" id="add-user-gender" placeholder="Male"
                                           onChange={e => this.setState({gender: e.target.value})}/>
                                    {/* TODO: get list from server and use select box */}
                                    Birth:
                                    <input type="text" className="form-control" id="add-user-birth" placeholder="yyyy"
                                           onChange={e => this.setState({birth: e.target.value})}/>
                                    <div className="modal-footer" style={{paddingBottom: "0"}}>
                                        <button type="button" className="btn btn-primary" data-dismiss="modal"
                                                onClick={this.handleSignUpClick}>
                                            Submit
                                        </button>
                                    </div>
                                </div>
                                <div className="tab-pane fade" id="login-tab" role="tabpanel">
                                    ID:
                                    <input type="text" className="form-control" id="login-user-id" placeholder="asdf"
                                           onChange={e => this.setState({userIdInput: e.target.value})}/>
                                    <div className="modal-footer" style={{paddingBottom: "0"}}>
                                        <button type="button" className="btn btn-primary" data-dismiss="modal"
                                                onClick={this.handleLoginClick}>
                                            Submit
                                        </button>
                                    </div>
                                </div>
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
                console.log("user present", data);
                this.setState({userId: data.userID});
            })
            .catch(() => {
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

    handleSignUpClick() {
        axios.post("http://localhost:5000/api/v1/users", {
            nationality: this.state.nationality,
            gender: this.state.gender,
            birth: Number(this.state.birth)
        }).then(({data}) => {
            this.setState({userId: data.id});
        });
    }

    handleLoginClick() {
        axios.post("http://localhost:5000/api/v1/users/presence", {id: this.state.userIdInput})
            .then(({data}) => {
                this.setState({userId: data.userID});
            })
            .catch(() => {
                // TODO: error handling
            });
    }
}

export default Header;