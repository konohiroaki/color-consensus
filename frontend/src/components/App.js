import React, {Component} from "react";
import Header from "./header/Header";
import Body from "./body/Body";
import LoginModal from "./common/LoginModal";

// TODO: do test for react app (jest?)
class App extends Component {

    constructor(props) {
        super(props);
        this.state = {};
        /* userId: undefined -> haven't checked user login status.
         *         null      -> user isn't logged in.
         *         "string"  -> user is logged in.
         */

        this.setUserId = this.setUserId.bind(this);
    }

    setUserId(userId) {
        this.setState({userId: userId});
    }

    render() {
        return (
            <div className="bg-dark text-light" style={{display: "flex", flexDirection: "column", height: "100%"}}>
                <LoginModal setUserId={this.setUserId} userId={this.state.userId}/>
                <Header style={{flex: "0 0 80px"}} userId={this.state.userId}/>
                <Body style={{flex: "1 1 auto"}} userId={this.state.userId}/>
            </div>
        );
    }
}

export default App;
