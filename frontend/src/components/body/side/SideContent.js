import React, {Component} from "react";
import axios from "axios";
import NewColorCard from "./NewColorCard";
import {isSameColor, isUndefinedColor} from "../../common/Utility";

class SideContent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            colorList: [],
            langFilter: "",
            nameFilter: "",
        };

        this.updateColorList = this.updateColorList.bind(this);
    }

    componentDidMount() {
        this.updateColorList().then(() => {
            // TODO: select random color in user's language?
            this.props.setTarget(this.state.colorList[0]);
        });
    }

    render() {
        console.log("rendering side content");

        return <div className={this.props.className} style={this.props.style}>
            <SearchBar langList={this.state.colorList.map(color => color.lang)}
                       langFilter={this.state.langFilter} langFilterSetter={e => this.setState({langFilter: e.target.value})}
                       nameFilter={this.state.nameFilter} nameFilterSetter={e => this.setState({nameFilter: e.target.value})}/>
            <Cards colorList={this.state.colorList} target={this.props.target} setTarget={this.props.setTarget}
                   langFilter={this.state.langFilter} nameFilter={this.state.nameFilter}
                   updateColorList={this.updateColorList}/>
        </div>;
    }

    updateColorList() {
        return axios.get(`${process.env.WEBAPI_HOST}/api/v1/colors/keys`).then(({data}) => {
            console.log("side content got color list from server: ", data);
            this.setState({colorList: data});
        });
    }
}

const SearchBar = props => (
    <div className="input-group rounded-0">
        <LangFilterSelector langList={props.langList}
                            langFilter={props.langFilter} langFilterSetter={props.langFilterSetter}/>
        <NameFilterInput nameFilter={props.nameFilter} nameFilterSetter={props.nameFilterSetter}/>
    </div>
);

const LangFilterSelector = ({langList, langFilter, langFilterSetter}) => {
    const langOptions = getLangList(langList)
        .map(lang => <option key={lang} value={lang}>{lang !== "" ? lang : "Language"}</option>);

    return <div className="input-group-prepend">
        <select className="custom-select rounded-0" value={langFilter} onChange={langFilterSetter}>
            {langOptions}
        </select>
    </div>;
};

const getLangList = langList => {
    return langList.reduce((acc, current) => {
        if (!acc.includes(current)) {
            acc.push(current);
        }
        return acc;
    }, [""]);
};

const NameFilterInput = ({nameFilter, nameFilterSetter}) => (
    <input type="text" className="form-control" value={nameFilter} onChange={nameFilterSetter}/>
);

const Cards = (props) => (
    <div style={{overflowY: "auto", height: "100%"}}>
        <TargetColorCard target={props.target}/>
        <SelectableColorCards colorList={props.colorList} target={props.target} setTarget={props.setTarget}
                              langFilter={props.langFilter} nameFilter={props.nameFilter}/>
        <NewColorCard updateColorList={props.updateColorList}/>
    </div>
);

const TargetColorCard = ({target}) => {
    if (target === undefined) {
        return null;
    }
    return <div className="d-block m-2 card btn bg-dark text-light border border border-primary">
        <div className="row">
            <div className="col-3 border-right border-secondary p-3">{target.lang}</div>
            <div className="col-9 p-3">{target.name}</div>
        </div>
    </div>;
};

const SelectableColorCards = (props) => {
    const selectableCards = props.colorList
        .filter(c => !isSameColor(c, props.target))
        .filter(c => isLangMatchingFilter(c.lang, props.langFilter))
        .filter(c => isNameMatchingFilter(c.name, props.nameFilter))
        .sort(colorComparator)
        .map(c => <ColorCard key={c.lang + ":" + c.name} color={c} setTarget={props.setTarget}/>);

    return <div>{selectableCards}</div>;
};

const isLangMatchingFilter = (lang, filter) => filter === "" || lang === filter;
const isNameMatchingFilter = (name, filter) => filter === "" || name.includes(filter.toLowerCase());
const colorComparator = (c1, c2) => c1.lang !== c2.lang ? (c1.lang > c2.lang ? 1 : -1) : (c1.name > c2.name ? 1 : -1);

const ColorCard = ({color, setTarget}) => (
    <div className="d-block m-2 card btn bg-dark text-light border border border-secondary" onClick={() => setTarget(color)}>
        <div className="row">
            <div className="col-3 border-right border-secondary p-3">{color.lang}</div>
            <div className="col-9 p-3">{color.name}</div>
        </div>
    </div>
);

export default SideContent;
