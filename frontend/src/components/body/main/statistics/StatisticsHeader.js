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
            "", this.props.baseColor.name,
            "", this.props.voteCount);
        return <div className="card bg-dark border border-secondary">
            <div className="card-body">
                <div className="row ml-0 mr-0">
                    <ColorCard baseColor={this.props.baseColor}/>
                    <StatisticsFilter/>
                    <VoteCounter voteCount={this.props.voteCount}/>
                </div>
            </div>
        </div>;
    }
}

const ColorCard = ({baseColor}) => (
    <div className="col-3 card bg-dark border border-secondary p-2 text-center">
        <div className="row">
            <span className="col-4 border-right border-secondary p-3">{baseColor.lang}</span>
            <span className="col-8 p-3">{baseColor.name}</span>
        </div>
    </div>
);

// TODO: complete select box impl
const StatisticsFilter = () => (
    <div className="col-7 input-group">
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

const VoteCounter = (props) => (
    <div className="col-2 card bg-dark border border-secondary p-2 text-center">
        Vote count
        <div className="font-weight-bold">{props.voteCount}</div>
    </div>
);

export default StatisticsHeader;
