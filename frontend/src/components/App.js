import React, {Component} from "react";
import {connect} from "react-redux";
import Header from "./header/Header";
import Body from "./body/Body";
import LoginModal from "./common/LoginModal";
import {actions as user} from "../ducks/user";

// TODO: do test for react app (jest?)
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
    }
}

const mapDispatchToProps = dispatch => ({
    verifyLoginState: () => dispatch(user.verifyLoginState()),
});

export default connect(null, mapDispatchToProps)(App);
