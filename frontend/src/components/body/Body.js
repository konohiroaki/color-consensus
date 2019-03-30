import React, {Component} from "react";
import MainContent from "./main/MainContent";
import SideContent from "./side/SideContent";
import {BrowserRouter as Router} from "react-router-dom";

class Body extends Component {

    render() {
        console.log("rendering [body]");
        return <Router>
            <div className="d-flex flex-row" style={this.props.style}>
                <MainContent style={{flex: "1 1 auto"}} loginModalRef={this.props.loginModalRef}/>
                <SideContent style={{flex: "0 0 auto"}} className="border-left border-secondary"/>
            </div>
        </Router>;
    }
}

export default Body;
