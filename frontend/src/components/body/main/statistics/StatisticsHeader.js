import React, {Component} from "react";
import {actions as statistics} from "../../../../modules/statistics";
import {connect} from "react-redux";

class StatisticsHeader extends Component {

    render() {
        console.log("rendering [statistics header]",
            "baseColor.code:", this.props.baseColor.code,
            "votes.length:", this.props.votes.length);
        return <div className="card bg-dark border border-secondary">
            <div className="card-body">
                <div className="row ml-0 mr-0">
                    <div className="col-10">
                        <StatisticsFilter votes={this.props.votes} nationalityFilter={this.props.nationalityFilter}
                                          setNationalityFilter={this.props.setNationalityFilter}/>
                        <StatisticsPercentile/>
                    </div>
                    <VoteCounter style="col-2" voteCount={this.props.votes.length}/>
                </div>
            </div>
        </div>;
    }
}

// TODO: complete select box impl
const StatisticsFilter = (props) => {
    const nationalities = props.votes
        .map(v => v.voter.nationality)
        .filter(distinct)
        .map(n => <option key={n} value={n}>{n}</option>);

    return <div>
        Filters
        <div className="input-group">
            <select className="custom-select" value={props.nationalityFilter}
                    onChange={e => props.setNationalityFilter(e.target.value)}>
                <option value="">Nationality</option>
                {nationalities}
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
    </div>;
};

// https://stackoverflow.com/a/14438954
const distinct = (value, index, self) => self.indexOf(value) === index;

const StatisticsPercentile = () => (
    <div>
        Percentile
        <input type="range" className="custom-range" min="0" max="100" step="10"/>
    </div>
);

const VoteCounter = ({style, voteCount}) => (
    <div className={style + " card bg-dark border border-secondary p-2 text-center my-auto"}>
        Vote count
        <div className="font-weight-bold">{voteCount}</div>
    </div>
);

const mapStateToProps = state => ({
    votes: state.statistics.votes,
    nationalityFilter: state.statistics.nationalityFilter,
});

const mapDispatchToProps = dispatch => ({
    setNationalityFilter: nationality => dispatch(statistics.setNationalityFilter(nationality)),
});

export default connect(mapStateToProps, mapDispatchToProps)(StatisticsHeader);
