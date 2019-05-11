import React, {Component} from "react";
import $ from "jquery";
import {connect} from "react-redux";
import {actions as user} from "../../modules/user";
import {toast} from "react-toastify";

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
                        nationalities={this.props.nationalities}
                        setNationalityInput={input => this.setState({nationality: input})}
                        setBirthInput={input => this.setState({birth: input})}
                        genders={this.props.genders}
                        setGenderInput={input => this.setState({gender: input})}
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
        this.props.signUp(this.state.nationality, this.state.birth, this.state.gender)
            .then(() => $("#signup-login-modal").modal("toggle"))
            .then(() => this.runAndResetCallback())
            .catch(({response}) => toast.warn(response.data.error.message));
    }

    handleLoginClick() {
        this.props.login(this.state.userIdInput)
            .then(() => $("#signup-login-modal").modal("toggle"))
            .then(() => this.runAndResetCallback())
            .catch(({response}) => toast.warn(response.data.error.message));
    }

    runAndResetCallback() {
        this.callback();
        this.callback = () => {};
    }
}

const ModalHeader = () => (
    <div className="modal-header pb-0">
        <ul className="nav nav-tabs border-bottom-0" role="tablist">
            <li className="nav-item text-light"><SignUpTabLink/></li>
            <li className="nav-item text-light"><LoginTabLink/></li>
        </ul>
    </div>
);

const SignUpTabLink = () => <div className="nav-link active" href="#signup-tab" data-toggle="tab" role="tab">Sign Up</div>;
const LoginTabLink = () => <div className="nav-link" href="#login-tab" data-toggle="tab" role="tab">Login</div>;

const ModalBody = props => (
    <div className="modal-body">
        <TabContents props={props}/>
    </div>
);

const TabContents = ({props}) => (
    <div className="tab-content">
        <SignUpTabPanel nationalities={props.nationalities} setNationalityInput={props.setNationalityInput}
                        setBirthInput={props.setBirthInput}
                        genders={props.genders} setGenderInput={props.setGenderInput}
                        handleSignUpClick={props.handleSignUpClick}/>
        <LoginTabPanel setUserIdInput={props.setUserIdInput} handleLoginClick={props.handleLoginClick}/>
    </div>
);

const SignUpTabPanel = props => (
    <div className="tab-pane fade show active" id="signup-tab" role="tabpanel">
        <SignUpNationalityInput nationalities={props.nationalities}
                                setNationalityInput={props.setNationalityInput}/>
        <SignUpBirthInput setBirthInput={props.setBirthInput}/>
        <SignUpGenderInput genders={props.genders} setGenderInput={props.setGenderInput}/>

        <SignUpButton handleSignUpClick={props.handleSignUpClick}/>
    </div>
);

const SignUpNationalityInput = props => {
    var nationalities = Object.keys(props.nationalities)
        .sort((a, b) => props.nationalities[a] > props.nationalities[b] ? 1 : -1)
        .map(n => <option key={n} value={n}>{props.nationalities[n]}</option>);
    nationalities.unshift(<option key="0" value="">Choose from dropdown</option>);

    return <div>
        <label className="mb-0">Nationality:</label>
        <select className="custom-select"
                onChange={e => props.setNationalityInput(e.target.value)}>
            {nationalities}
        </select>
    </div>;
};

const SignUpBirthInput = ({setBirthInput}) => (
    <div>
        <label className="mb-0">Birth Year:</label>
        <input type="number" min="1900" max={new Date().getFullYear()} required
               className="form-control" placeholder="1990"
               onChange={e => setBirthInput(e.target.value)}/>
    </div>
);

const SignUpGenderInput = props => {
    var genders = props.genders
        .map(g => <option key={g} value={g}>{g}</option>);
    genders.unshift(<option key="0" value="">Choose from dropdown</option>);

    return <div>
        <label className="mb-0">Gender:</label>
        <select className="custom-select"
                onChange={e => props.setGenderInput(e.target.value)}>
            {genders}
        </select>
    </div>;
};

const SignUpButton = ({handleSignUpClick}) => (
    <div className="col-3 ml-auto pt-3">
        <button type="button" className="btn btn-primary" onClick={handleSignUpClick}>
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
        <label className="mb-0">ID:</label>
        <input type="text" className="form-control" placeholder="00000000-0000-0000-0000-000000000000"
               onChange={e => setUserIdInput(e.target.value)}/>
    </div>
);

const LoginButton = ({handleLoginClick}) => (
    <div className="col-3 ml-auto pt-3">
        <button type="button" className="btn btn-primary" onClick={handleLoginClick}>
            Submit
        </button>
    </div>
);

const mapStateToProps = state => ({
    userId: state.user.id,
    nationalities: state.nationality.nationalities,
    genders: state.gender.genders,
});

const mapDispatchToProps = dispatch => ({
    login: (id) => dispatch(user.login(id)),
    signUp: (nationality, birth, gender) => dispatch(user.signUp(nationality, birth, gender))
});

export default connect(mapStateToProps, mapDispatchToProps, null, {forwardRef: true})(LoginModal);
