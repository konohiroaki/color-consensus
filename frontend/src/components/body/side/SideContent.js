import React, {Component} from "react";
import axios from "axios";
import AddColorCard from "./AddColorCard";

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
            <Cards colorList={this.state.colorList} setTarget={this.props.setTarget}
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
        <ColorCards colorList={props.colorList} setTarget={props.setTarget}
                    langFilter={props.langFilter} nameFilter={props.nameFilter}/>
        <AddColorCard updateColorList={props.updateColorList}/>
    </div>
);

const ColorCards = (props) => {
    const colorCards = props.colorList
        .filter(color => props.nameFilter === "" || color.name.includes(props.nameFilter.toLowerCase()))
        .filter(color => props.langFilter === "" || color.lang === props.langFilter)
        // TODO: sort on server side.
        // ascending order for lang -> name
        .sort((a, b) => a.lang === b.lang ? a.name - b.name : a.lang - b.lang)
        .map(color => <ColorCard key={color.lang + ":" + color.name} color={color} setTarget={props.setTarget}/>);

    return <div>{colorCards}</div>;
};

const ColorCard = ({color, setTarget}) => (
    <div className="d-block m-2 card btn bg-dark text-light border border border-secondary" onClick={() => setTarget(color)}>
        <div className="row">
            <div className="col-3 border-right border-secondary p-3">{color.lang}</div>
            <div className="col-9 p-3">{color.name}</div>
        </div>
    </div>
);

export default SideContent;
