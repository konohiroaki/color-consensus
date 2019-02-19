import React, {Component} from "react";
import MainContent from "./main/MainContent";
import SideContent from "./side/SideContent";

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
        console.log("rendering body");
        return (
            <div style={Object.assign({display: "flex", flexDirection: "row"}, this.props.style)}>
                <MainContent style={{flex: "1 1 auto"}} target={this.state.target}/>
                <SideContent style={{flex: "0 0 auto"}} className="border-left border-secondary" setTarget={this.setTarget}/>
            </div>
        );
    }
}

export default Body;