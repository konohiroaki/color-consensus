import React, {Component} from "react";

class StatisticsHeader extends Component {

    shouldComponentUpdate(nextProps) {
        if (JSON.stringify(this.props.baseColor) === JSON.stringify(nextProps.baseColor)
            && this.props.voteCount === nextProps.voteCount) {
            return false;
        }
        return true;
    }

    render() {
        console.log("rendering [statistics header]",
            "baseColor.code:", this.props.baseColor.code,
            "voteCount:", this.props.voteCount);
        return <div className="card bg-dark border border-secondary">
            <div className="card-body">
                <div className="row ml-0 mr-0">
                    <div className="col-10">
                        <StatisticsFilter/>
                        <StatisticsPercentile/>
                    </div>
                    <VoteCounter style="col-2" voteCount={this.props.voteCount}/>
                </div>
            </div>
        </div>;
    }
}

// TODO: complete select box impl
const StatisticsFilter = () => (
    <div>
        Filters
        <div className="input-group">
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
);

const StatisticsPercentile = () => (
    <div>
        Percentile
        <input type="range" className="custom-range" min="0" max="100" step="10"/>
    </div>
);

const VoteCounter = ({style, voteCount}) => (
    <div className={style + " card bg-dark border border-secondary p-2 text-center"}>
        Vote count
        <div className="font-weight-bold">{voteCount}</div>
    </div>
);

export default StatisticsHeader;
