import React, {Component} from "react";
import axios from "axios";
import $ from "jquery";

class LoginModal extends Component {

    openLoginModal(callback) {
        $("#signup-login-modal").modal();
        if (callback !== undefined) {
            this.callback = callback;
        } else {
            console.log("callback is undefined");
        }
    }

    constructor(props) {
        super(props);
        this.state = {};

        this.callback = () => {};

        this.handleSignUpClick = this.handleSignUpClick.bind(this);
        this.handleLoginClick = this.handleLoginClick.bind(this);
        this.runAndResetCallback = this.runAndResetCallback.bind(this);
    }

    runAndResetCallback() {
        this.callback();
        this.callback = () => {};
    }

    render() {
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
        axios.get(`${process.env.WEBAPI_HOST}/api/v1/users/presence`)
            .then(({data}) => this.props.setUserId(data.userID))
            .catch(() => this.props.setUserId(null));
    }

    handleSignUpClick() {
        axios.post(`${process.env.WEBAPI_HOST}/api/v1/users`, {
            nationality: this.state.nationality,
            gender: this.state.gender,
            birth: Number(this.state.birth)
        }).then(({data}) => {
            this.props.setUserId(data.id);
            this.runAndResetCallback();
        });
    }

    handleLoginClick() {
        axios.post(`${process.env.WEBAPI_HOST}/api/v1/users/presence`, {id: this.state.userIdInput})
            .then(() => {
                this.props.setUserId(this.state.userIdInput);
                this.runAndResetCallback();
            })
            .catch(() => {
                // TODO: error handling
            });
    }
}

export default LoginModal;
