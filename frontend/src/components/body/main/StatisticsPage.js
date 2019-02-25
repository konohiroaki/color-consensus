import React, {Component} from "react";

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
                        <button className="btn btn-secondary m-3" onClick={() => this.props.history.push("/")}>
                            Back to voting
                        </button>
                    </div>
                </div>

                {/* TODO: add content here. */}
                <div>Hey!</div>
            </div>
        );
    }
}

export default StatisticsPage;
