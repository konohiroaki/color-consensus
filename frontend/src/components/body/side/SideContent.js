import React, {Component} from "react";
import AddColorCard from "./AddColorCard";
import {actions as searchBar} from "../../../modules/searchBar";
import {actions as board} from "../../../modules/board";
import {connect} from "react-redux";

class SideContent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            langFilter: "",
            nameFilter: "",
        };
    }

    componentDidMount() {
        this.props.fetchColors();
    }

    componentDidUpdate() {
        if (this.props.baseColor === null) {
            this.props.setBaseColor();
        }
    }

    render() {
        console.log("rendering [side content]",
            "list.length:", this.props.colorList.length,
            "base:", this.props.baseColor !== null ? this.props.baseColor.code : null);
        return <div className={this.props.className + " d-flex flex-column"}
                    style={this.props.style}>
            <SearchBar colorList={this.props.colorList}
                       langFilter={this.state.langFilter} setLangFilter={e => this.setState({langFilter: e})}
                       nameFilter={this.state.nameFilter} setNameFilter={e => this.setState({nameFilter: e})}/>
            <Cards colorList={this.props.colorList} baseColor={this.props.baseColor}
                   langFilter={this.state.langFilter} nameFilter={this.state.nameFilter}
                   setBaseColor={this.props.setBaseColor}/>
        </div>;
    }
}

const SearchBar = (props) => {
    return <div className="input-group rounded-0">
        <LangFilterSelector colorList={props.colorList}
                            langFilter={props.langFilter} setLangFilter={props.setLangFilter}/>
        <NameFilterInput nameFilter={props.nameFilter} setNameFilter={props.setNameFilter}/>
    </div>;
};

const LangFilterSelector = ({colorList, langFilter, setLangFilter}) => {
    const langOptions = getLangList(colorList)
        .map(lang => <option key={lang} value={lang}>{lang !== "" ? lang : "Language"}</option>);

    return <div className="input-group-prepend">
        <select className="custom-select rounded-0" value={langFilter}
                onChange={e => setLangFilter(e.target.value)}>
            {langOptions}
        </select>
    </div>;
};

const NameFilterInput = ({nameFilter, setNameFilter}) => {
    return <input type="text" className="form-control" value={nameFilter}
                  onChange={e => setNameFilter(e.target.value)}/>;
};

const Cards = (props) => {
    return <div style={{height: "100%", overflowY: "auto"}}>
        <BaseColorCard baseColor={props.baseColor}/>
        <SelectableColorCards colorList={props.colorList} baseColor={props.baseColor}
                              langFilter={props.langFilter} nameFilter={props.nameFilter}
                              setBaseColor={props.setBaseColor}/>
        <AddColorCard/>
    </div>;
};

const BaseColorCard = ({baseColor}) => {
    if (baseColor === null) {
        return null;
    }
    return <div className="d-block m-2 card btn bg-dark text-light border border border-primary">
        <div className="row">
            <div className="col-3 border-right border-secondary p-3">{baseColor.lang}</div>
            <div className="col-9 p-3">{baseColor.name}</div>
        </div>
    </div>;
};

const SelectableColorCards = ({colorList, baseColor, langFilter, nameFilter, setBaseColor}) => {
    const selectableCards = colorList
        .filter(c => baseColor !== null && !isSameColor(c, baseColor))
        .filter(c => isLangMatchingFilter(c.lang, langFilter))
        .filter(c => isNameMatchingFilter(c.name, nameFilter))
        .sort(colorComparator)
        .map(c => <ColorCard key={c.lang + ":" + c.name} color={c} setBaseColor={setBaseColor}/>);

    return selectableCards.length !== 0 ? <div>{selectableCards}</div> : null;
};

const ColorCard = ({color, setBaseColor}) => {
    return <div className="d-block m-2 card btn bg-dark text-light border border border-secondary"
                onClick={() => setBaseColor(color)}>
        <div className="row">
            <div className="col-3 border-right border-secondary p-3">{color.lang}</div>
            <div className="col-9 p-3">{color.name}</div>
        </div>
    </div>;
};

const getLangList = colorList => colorList.map(color => color.lang)
    .reduce((acc, current) => {
        if (!acc.includes(current)) {
            acc.push(current);
        }
        return acc;
    }, [""]);

const isLangMatchingFilter = (lang, filter) => filter === "" || lang === filter;
const isNameMatchingFilter = (name, filter) => filter === "" || name.includes(filter.toLowerCase());
const colorComparator = (c1, c2) => c1.lang !== c2.lang ? (c1.lang > c2.lang ? 1 : -1) : (c1.name > c2.name ? 1 : -1);
const isSameColor = (c1, c2) => c1 !== undefined && c2 !== undefined
                                && c1.lang === c2.lang && c1.name === c2.name && c1.code === c2.code;

const mapStateToProps = state => ({
    colorList: state.searchBar.colors,
    baseColor: state.board.baseColor,
});

const mapDispatchToProps = dispatch => ({
    fetchColors: () => dispatch(searchBar.fetchColors()),
    setBaseColor: color => dispatch(board.setBaseColor(color)),
});

export default connect(mapStateToProps, mapDispatchToProps)(SideContent);
