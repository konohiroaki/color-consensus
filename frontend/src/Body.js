import React, {Component} from "react";
import MainContent from "./MainContent";
import SideContent from "./SideContent";

class Body extends Component {

    constructor(props) {
        super(props);
        this.state = {
            target: {}
        };
        this.setTarget = this.setTarget.bind(this);
    }

    setTarget(target) {
        console.log("body got target: ",target);
        this.setState({target: target});
    }

    // TODO: route to statistics page (https://reacttraining.com/react-router/)
    render() {
        return (
            <div style={Object.assign({display: "flex", flexDirection: "row"}, this.props.style)}>
                <MainContent style={{flex: "1 1 auto"}} target={this.state.target}/>
                <SideContent style={{flex: "0 0 auto"}} className="border-left border-secondary" setTarget={this.setTarget}/>
            </div>
        );
    }
}

export default Body;