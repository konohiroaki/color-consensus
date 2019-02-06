import React, {Component} from "react";
import MainContent from "./MainContent";
import SideBar from "./SideBar";

class Body extends Component {

    render() {
        return (
            <div style={Object.assign(this.props.style, {display: "flex", flexDirection: "row"})}>
                <MainContent/>
                <SideBar className="border-left border-secondary"/>
            </div>
        );
    }
}

export default Body;