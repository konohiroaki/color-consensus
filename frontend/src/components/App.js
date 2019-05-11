import React, {Component} from "react";
import {connect} from "react-redux";
import Header from "./header/Header";
import Body from "./body/Body";
import LoginModal from "./common/LoginModal";
import {actions as user} from "../modules/user";
import {actions as language} from "../modules/language";
import {actions as nationality} from "../modules/nationality";
import {actions as gender} from "../modules/gender";

class App extends Component {

    constructor(props) {
        super(props);
        this.loginModalRef = React.createRef();
    }

    render() {
        return <div className="d-flex flex-column bg-dark text-light" style={{height: "100%"}}>
            <LoginModal ref={this.loginModalRef}/>

            <Header style={{flex: "0 0 80px"}} loginModalRef={this.loginModalRef}/>
            <Body style={{flex: "1 1 auto"}} loginModalRef={this.loginModalRef}/>
        </div>;
    }

    componentDidMount() {
        this.props.verifyLoginState();
        this.props.setLanguages();
        this.props.setNationalities();
        this.props.setGenders();
    }
}

const mapDispatchToProps = dispatch => ({
    verifyLoginState: () => dispatch(user.verifyLoginState()),
    setLanguages: () => dispatch(language.setLanguages()),
    setNationalities: () => dispatch(nationality.setNationalities()),
    setGenders: () => dispatch(gender.setGenders()),
});

export default connect(null, mapDispatchToProps)(App);
