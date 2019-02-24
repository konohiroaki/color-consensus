import React, {Component} from "react";
import {Link} from "react-router-dom";

class StatisticsPage extends Component {
    render() {
        console.log("rendering statistics page");
        if (this.props.target === undefined) {
            return null;
        }

        return (
            <div>
                <div className="row">
                    <div className="mr-auto ml-5">
                        <p>Language: {this.props.target.lang}</p>
                        <p>Color Name: {this.props.target.name}</p>
                    </div>

                    <div className="ml-auto">
                        <Link to={"/"}>
                            <button className="btn btn-secondary m-3">Back to voting</button>
                        </Link>
                    </div>
                </div>

                {/* TODO: add content here. */}
                <div>Hey!</div>
            </div>
        );
    }
}

export default StatisticsPage;