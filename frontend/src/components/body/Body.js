import React, {Component} from "react";
import MainContent from "./main/MainContent";
import SideContent from "./side/SideContent";
import {BrowserRouter as Router} from "react-router-dom";

class Body extends Component {

    constructor(props) {
        super(props);
        this.state = {};
        this.setTarget = this.setTarget.bind(this);
    }

    setTarget(target) {
        this.setState({target: target});
    }

    render() {
        console.log("rendering body", this.props.userId, this.state.target);
        return <Router>
            <div className="d-flex flex-row" style={this.props.style}>
                <MainContent style={{flex: "1 1 auto"}} userId={this.props.userId} target={this.state.target}
                             loginModalRef={this.props.loginModalRef}/>
                <SideContent style={{flex: "0 0 auto"}} className="border-left border-secondary"
                             target={this.state.target} setTarget={this.setTarget}/>
            </div>
        </Router>;
    }
}

export default Body;
