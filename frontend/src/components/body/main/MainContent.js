import React, {Component} from "react";
import {Route, Switch} from "react-router-dom";
import VotingPage from "./vote/VotingPage";
import StatisticsPage from "./statistics/StatisticsPage";
import {Slide, ToastContainer} from "react-toastify";
import "react-toastify/dist/ReactToastify.min.css";

class MainContent extends Component {

    render() {
        console.log("rendering [main content]");

        return <div className="container-fluid"
                    style={Object.assign({overflowY: "auto", height: "100%"}, this.props.style)}>
            <ToastContainer
                position="top-center"
                autoClose={5000000}
                hideProgressBar={false}
                newestOnTop
                closeOnClick
                pauseOnVisibilityChange
                pauseOnHover
                transition={Slide}
            />
            <Switch>
                <Route exact path={"/"} render={({history}) => (
                    <VotingPage history={history} loginModalRef={this.props.loginModalRef}/>
                )}/>
                <Route path={"/statistics"} render={({history}) => (
                    <StatisticsPage history={history}/>
                )}/>
            </Switch>
        </div>;
    }
}

export default MainContent;
