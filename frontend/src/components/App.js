import React, {Component, createRef} from "react";
import Header from "./header/Header";
import Body from "./body/Body";
import LoginModal from "./common/LoginModal";

// TODO: do updateSelectedState for react app (jest?)
class App extends Component {

    constructor(props) {
        super(props);
        this.state = {};
        /* userId: undefined -> haven't checked user login status.
         *         null      -> user isn't logged in.
         *         "string"  -> user is logged in.
         */

        this.loginModalRef = React.createRef();

        this.setUserId = this.setUserId.bind(this);
    }

    setUserId(userId) {
        this.setState({userId: userId});
    }

    render() {
        const loginModalRef = this.loginModalRef.current;

        return (
            <div className="bg-dark text-light" style={{display: "flex", flexDirection: "column", height: "100%"}}>
                <LoginModal ref={this.loginModalRef}
                            setUserId={(userId) => this.setState({userId: userId})}
                            userId={this.state.userId}/>
                <Header style={{flex: "0 0 80px"}} userId={this.state.userId} loginModalRef={loginModalRef}/>
                <Body style={{flex: "1 1 auto"}} userId={this.state.userId} loginModalRef={loginModalRef}/>
            </div>
        );
    }
}

export default App;
