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
                        <StatisticsFilter votes={this.props.votes}
                                          nationalityFilter={this.props.nationalityFilter}
                                          setNationalityFilter={this.props.setNationalityFilter}
                                          ageGroupFilter={this.props.ageGroupFilter}
                                          setAgeGroupFilter={this.props.setAgeGroupFilter}
                                          genderFilter={this.props.genderFilter}
                                          setGenderFilter={this.props.setGenderFilter}/>
                        <StatisticsPercentile percentile={this.props.percentile} setPercentile={this.props.setPercentile}/>
                    </div>
                    <VoteCounter style="col-2" voteCount={this.props.votes.length}/>
                </div>
            </div>
        </div>;
    }
}

const StatisticsFilter = (props) => {
    const nationalities = props.votes
        .map(v => v.voter.nationality)
        .filter(distinct)
        .sort()
        .map(n => <option key={n} value={n}>{n}</option>);
    const ageGroups = props.votes
        .map(v => v.voter.ageGroup)
        .filter(distinct)
        .sort()
        .map(n => <option key={n} value={n}>{n + " ~ " + (n + 9)}</option>);
    const genders = props.votes
        .map(v => v.voter.gender)
        .filter(distinct)
        .sort()
        .map(n => <option key={n} value={n}>{n}</option>);

    return <div>
        Filters:
        <div className="input-group">
            <select className="custom-select" value={props.nationalityFilter}
                    onChange={e => props.setNationalityFilter(e.target.value)}>
                <option value="">Nationality</option>
                {nationalities}
            </select>
            <select className="custom-select" value={props.ageGroupFilter}
                    onChange={e => props.setAgeGroupFilter(e.target.value)}>
                <option value="">Age Group</option>
                {ageGroups}
            </select>
            <select className="custom-select" value={props.genderFilter}
                    onChange={e => props.setGenderFilter(e.target.value)}>
                <option value="">Gender</option>
                {genders}
            </select>
        </div>
    </div>;
};

// https://stackoverflow.com/a/14438954
const distinct = (value, index, self) => self.indexOf(value) === index;

const StatisticsPercentile = (props) => {
    let percentile = parseInt(props.percentile);
    if (percentile === 0) {
        percentile = "At least one person voted";
    } else if (percentile === 100) {
        percentile = "Everyone voted";
    } else {
        percentile += "%~ voted";
    }
    return <div>
        Percentile: [{percentile}]
        <input type="range" className="custom-range" min="0" max="100" step="10"
               onChange={e => {
                   props.setPercentile(e.target.value);
               }}/>
    </div>;
};

const VoteCounter = ({style, voteCount}) => (
    <div className={style + " card bg-dark border border-secondary p-2 text-center my-auto"}>
        Vote count
        <div className="font-weight-bold">{voteCount}</div>
    </div>
);

const mapStateToProps = state => ({
    votes: state.statistics.votes,
    nationalityFilter: state.statistics.nationalityFilter,
    ageGroupFilter: state.statistics.ageGroupFilter,
    genderFilter: state.statistics.genderFilter,
    percentile: state.statistics.percentile,
});

const mapDispatchToProps = dispatch => ({
    setNationalityFilter: nationality => dispatch(statistics.setNationalityFilter(nationality)),
    setAgeGroupFilter: ageGroup => dispatch(statistics.setAgeGroupFilter(ageGroup)),
    setGenderFilter: gender => dispatch(statistics.setGenderFilter(gender)),
    setPercentile: percentile => dispatch(statistics.setPercentile(percentile))
});

export default connect(mapStateToProps, mapDispatchToProps)(StatisticsHeader);
