import React, {Component} from "react";

class StatisticsPage extends Component {
    render() {
        console.log("rendering statistics page");
        if (this.props.target === undefined) {
            return null;
        }

        return (
            <div>
                <div className="card bg-dark border border-secondary">
                    <div className="card-body">
                        <div className="row ml-0 mr-0">
                            <div className="col-3 card bg-dark border border-secondary p-2 text-center">
                                <div className="row">
                                    <span className="col-4 border-right border-secondary p-3">{this.props.target.lang}</span>
                                    <span className="col-8 p-3">{this.props.target.name}</span>
                                </div>
                            </div>
                            {/* TODO: complete select box impl */}
                            <div className="col-9 input-group">
                                <select className="custom-select">
                                    <option selected>Nationality</option>
                                    <option value="3">Others</option>
                                </select>
                                <select className="custom-select">
                                    <option selected>Age Group</option>
                                    <option value="3">Others</option>
                                </select>
                                <select className="custom-select">
                                    <option selected>Gender</option>
                                    <option value="1">Male</option>
                                    <option value="2">Female</option>
                                    <option value="3">Others</option>
                                </select>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="row">
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
