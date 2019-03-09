import React, {Component} from "react";

class StatisticsContentHeader extends Component {

    render() {
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
                                <select className="custom-select" defaultValue="">
                                    <option value="">Nationality</option>
                                    <option value="3">Others</option>
                                </select>
                                <select className="custom-select" defaultValue="">
                                    <option value="">Age Group</option>
                                    <option value="3">Others</option>
                                </select>
                                <select className="custom-select" defaultValue="">
                                    <option value="">Gender</option>
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
            </div>
        );
    }
}

export default StatisticsContentHeader;
