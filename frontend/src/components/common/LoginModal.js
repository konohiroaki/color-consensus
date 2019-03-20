import React, {Component} from "react";
import axios from "axios";
import $ from "jquery";

class LoginModal extends Component {

    constructor(props) {
        super(props);
        this.state = {};

        this.callback = () => {};

        this.handleSignUpClick = this.handleSignUpClick.bind(this);
        this.handleLoginClick = this.handleLoginClick.bind(this);
        this.runAndResetCallback = this.runAndResetCallback.bind(this);
    }

    render() {
        return <div className="modal fade" tabIndex="-1" id="signup-login-modal">
            <div className="modal-dialog modal-dialog-centered">
                <div className="modal-content bg-dark">
                    <ModalHeader/>
                    <ModalBody
                        setNationalityInput={input => this.setState({nationality: input})}
                        setGenderInput={input => this.setState({gender: input})}
                        setBirthInput={input => this.setState({birth: input})}
                        handleSignUpClick={this.handleSignUpClick}
                        setUserIdInput={input => this.setState({userIdInput: input})}
                        handleLoginClick={this.handleLoginClick}/>
                </div>
            </div>
        </div>;
    }

    componentDidMount() {
        axios.get(`${process.env.WEBAPI_HOST}/api/v1/users/presence`)
            .then(({data}) => this.props.setUserId(data.userID))
            .catch(() => this.props.setUserId(null));
    }

    openLoginModal(callback) {
        if (callback !== undefined) {
            this.callback = callback;
        }
        $("#signup-login-modal").modal();
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

    runAndResetCallback() {
        this.callback();
        this.callback = () => {};
    }
}

const ModalHeader = () => (
    <div className="modal-header pb-0">
        <ul className="nav nav-tabs border-bottom-0" role="tablist">
            <li className="nav-item text-light"><SingUpTabLink/></li>
            <li className="nav-item text-light"><LoginTabLink/></li>
        </ul>
    </div>
);

const ModalBody = props => (
    <div className="modal-body">
        <TabContents props={props}/>
    </div>
);

const SingUpTabLink = () => <div className="nav-link active" href="#signup-tab" data-toggle="tab" role="tab">Sign Up</div>;
const LoginTabLink = () => <div className="nav-link" href="#login-tab" data-toggle="tab" role="tab">Login</div>;

const TabContents = ({props}) => (
    <div className="tab-content">
        <SingUpTabPanel setNationalityInput={props.setNationalityInput} setGenderInput={props.setGenderInput}
                        setBirthInput={props.setBirthInput} handleSignUpClick={props.handleSignUpClick}/>
        <LoginTabPanel setUserIdInput={props.setUserIdInput} handleLoginClick={props.handleLoginClick}/>
    </div>
);

const SingUpTabPanel = props => (
    <div className="tab-pane fade show active" id="signup-tab" role="tabpanel">
        {/* TODO: get list from server and use select box */}
        <SignUpNationalityInput setNationalityInput={props.setNationalityInput}/>
        {/* TODO: get list from server and use select box */}
        <SignUpGenderInput setGenderInput={props.setGenderInput}/>
        {/* TODO: get list from server and use select box */}
        <SignUpBirthInput setBirthInput={props.setBirthInput}/>

        <SignUpButton handleSignUpClick={props.handleSignUpClick}/>
    </div>
);

const SignUpNationalityInput = ({setNationalityInput}) => (
    <div>
        <label className="mb-0">Nationality:</label>
        <input type="text" className="form-control" placeholder="ex) Japan"
               onChange={e => setNationalityInput(e.target.value)}/>
    </div>
);

const SignUpGenderInput = ({setGenderInput}) => (
    <div>
        <label className="mb-0">Gender:</label>
        <input type="text" className="form-control" placeholder="ex) Male"
               onChange={e => setGenderInput(e.target.value)}/>
    </div>
);

const SignUpBirthInput = ({setBirthInput}) => (
    <div>
        <label className="mb-0">Birth:</label>
        <input type="text" className="form-control" placeholder="ex) 1990"
               onChange={e => setBirthInput(e.target.value)}/>
    </div>
);

const SignUpButton = ({handleSignUpClick}) => (
    <div className="col-3 ml-auto pt-3">
        <button type="button" className="btn btn-primary" data-dismiss="modal" onClick={handleSignUpClick}>
            Submit
        </button>
    </div>
);

const LoginTabPanel = ({setUserIdInput, handleLoginClick}) => (
    <div className="tab-pane fade" id="login-tab" role="tabpanel">
        <LoginIdInput setUserIdInput={setUserIdInput}/>
        <LoginButton handleLoginClick={handleLoginClick}/>
    </div>
);

const LoginIdInput = ({setUserIdInput}) => (
    <div>
        <label>ID:</label>
        <input type="text" className="form-control" placeholder="00000000-0000-0000-0000-000000000000"
               onChange={e => setUserIdInput(e.target.value)}/>
    </div>
);

const LoginButton = ({handleLoginClick}) => (
    <div className="col-3 ml-auto pt-3">
        <button type="button" className="btn btn-primary" data-dismiss="modal" onClick={handleLoginClick}>
            Submit
        </button>
    </div>
);

export default LoginModal;
