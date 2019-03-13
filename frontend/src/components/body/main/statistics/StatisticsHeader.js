import React, {Component} from "react";

class StatisticsHeader extends Component {

    render() {
        return <div className="card bg-dark border border-secondary">
            <div className="card-body">
                <div className="row ml-0 mr-0">
                    <ColorCard target={this.props.target}/>
                    <StatisticsFilter/>
                </div>
            </div>
        </div>;
    }
}

const ColorCard = ({target}) => (
    <div className="col-3 card bg-dark border border-secondary p-2 text-center">
        <div className="row">
            <span className="col-4 border-right border-secondary p-3">{target.lang}</span>
            <span className="col-8 p-3">{target.name}</span>
        </div>
    </div>
);

// TODO: complete select box impl
const StatisticsFilter = () => (
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
);

export default StatisticsHeader;
