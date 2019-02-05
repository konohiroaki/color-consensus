import React, {Component} from "react";

import SideBar from "./SideBar";
import MainContent from "./MainContent";

class App extends Component {

    // TODO: route to statistics page (https://reacttraining.com/react-router/)
    render() {
        return (
            <div className="bg-dark text-light" style={{display: "flex", flexDirection: "column", height: "100%"}}>
                <Header style={{flex: "0 0 80px"}}/>
                <div style={{flex: "1 1 auto", display: "flex", flexDirection: "row"}}>
                    <MainContent/>
                    <SideBar className="border-left border-secondary"/>
                </div>
            </div>
        );
    }
}

class Header extends Component {
    render() {
        return (
            <nav className="navbar navbar-dark bg-dark border-bottom border-secondary" style={this.props.style}>
                <a className="navbar-brand" href="/">Color Consensus</a>
            </nav>
        );
    }
}

export default App;
