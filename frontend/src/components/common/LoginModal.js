import React, {Component} from "react";
import $ from "jquery";
import {connect} from "react-redux";
import {actions as user} from "../../modules/user/user";

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

    openLoginModal(callback) {
        if (callback !== undefined) {
            this.callback = callback;
        }
        $("#signup-login-modal").modal();
    }

    handleSignUpClick() {
        this.props.signUp(this.state.nationality, this.state.gender, this.state.birth)
            .then(() => this.runAndResetCallback());
    }

    handleLoginClick() {
        this.props.login(this.state.userIdInput)
            .then(() => this.runAndResetCallback());
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

const mapStateToProps = state => ({
    userId: state.user.id,
});

const mapDispatchToProps = dispatch => ({
    login: (id) => dispatch(user.login(id)),
    signUp: (nationality, gender, birth) => dispatch(user.signUp(nationality, gender, birth))
});

export default connect(mapStateToProps, mapDispatchToProps, null, {forwardRef: true})(LoginModal);
