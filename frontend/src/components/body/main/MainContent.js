import React, {Component} from "react";
import {Route} from "react-router-dom";
import VotingPage from "./VotingPage";
import StatisticsPage from "./StatisticsPage";

class MainContent extends Component {

    render() {
        console.log("rendering main content");

        return (
            <div className="container-fluid pt-3" style={Object.assign({overflowY: "auto"}, this.props.style)}>
                <Route exact path={"/"} render={() => (
                    <VotingPage userId={this.props.userId} target={this.props.target}/>
                )}/>
                <Route path={"/statistics"} render={() => (
                    <StatisticsPage target={this.props.target}/>
                )}/>
            </div>
        );
    }
}

export default MainContent;