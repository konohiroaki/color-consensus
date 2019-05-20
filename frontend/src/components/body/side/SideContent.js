import React, {Component} from "react";
import AddColorCard from "./AddColorCard";
import {actions as searchBar} from "../../../modules/searchBar";
import {actions as board} from "../../../modules/board";
import {connect} from "react-redux";

class SideContent extends Component {

    constructor(props) {
        super(props);
        this.state = {
            categoryFilter: "",
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
                       categoryFilter={this.state.categoryFilter} setCategoryFilter={e => this.setState({categoryFilter: e})}
                       nameFilter={this.state.nameFilter} setNameFilter={e => this.setState({nameFilter: e})}/>
            <Cards colorList={this.props.colorList} baseColor={this.props.baseColor}
                   categoryFilter={this.state.categoryFilter} nameFilter={this.state.nameFilter}
                   setBaseColor={this.props.setBaseColor}/>
        </div>;
    }
}

const SearchBar = (props) => {
    return <div className="input-group rounded-0">
        <CategoryFilterSelector colorList={props.colorList}
                            categoryFilter={props.categoryFilter} setCategoryFilter={props.setCategoryFilter}/>
        <NameFilterInput nameFilter={props.nameFilter} setNameFilter={props.setNameFilter}/>
    </div>;
};

const CategoryFilterSelector = ({colorList, categoryFilter, setCategoryFilter}) => {
    const categoryOptions = getCategoryList(colorList)
        .map(category => <option key={category} value={category}>{category !== "" ? category : "Category"}</option>);

    return <div className="input-group-prepend">
        <select className="custom-select rounded-0" value={categoryFilter}
                onChange={e => setCategoryFilter(e.target.value)}>
            {categoryOptions}
        </select>
    </div>;
};

const NameFilterInput = ({nameFilter, setNameFilter}) => {
    return <input type="text" className="form-control" value={nameFilter}
                  onChange={e => setNameFilter(e.target.value)}/>;
};

const Cards = (props) => {
    return <div style={{height: "100%", overflowY: "auto"}}>
        <ColorCard color={props.baseColor} borderColor="border-primary" setBaseColor={{}}/>
        <SelectableColorCards colorList={props.colorList} baseColor={props.baseColor}
                              categoryFilter={props.categoryFilter} nameFilter={props.nameFilter}
                              setBaseColor={props.setBaseColor}/>
        <AddColorCard/>
    </div>;
};

const SelectableColorCards = ({colorList, baseColor, categoryFilter, nameFilter, setBaseColor}) => {
    const selectableCards = colorList
        .filter(c => baseColor !== null && !isSameColor(c, baseColor))
        .filter(c => isCategoryMatchingFilter(c.category, categoryFilter))
        .filter(c => isNameMatchingFilter(c.name, nameFilter))
        .sort(colorComparator)
        .map(c => <ColorCard key={c.category + ":" + c.name} color={c} borderColor="border-secondary" setBaseColor={setBaseColor}/>);

    return selectableCards.length !== 0 ? <div>{selectableCards}</div> : null;
};

const ColorCard = ({color, borderColor, setBaseColor}) => {
    if (color === null) {
        return null;
    }
    return <div className={"d-block m-2 card btn bg-dark text-light border border " + borderColor}
                onClick={() => setBaseColor(color)}>
        <div className="row">
            <div className="col-6 border-right border-secondary p-3 m-auto">{color.category}</div>
            <div className="col-6 p-3 m-auto">{color.name}</div>
        </div>
    </div>;
};

const getCategoryList = colorList => colorList.map(color => color.category)
    .reduce((acc, current) => {
        if (!acc.includes(current)) {
            acc.push(current);
        }
        return acc;
    }, [""]);

const isCategoryMatchingFilter = (category, filter) => filter === "" || category === filter;
const isNameMatchingFilter = (name, filter) => filter === "" || name.includes(filter.toLowerCase());
const colorComparator = (c1, c2) => c1.category !== c2.category ? (c1.category > c2.category ? 1 : -1) : (c1.name > c2.name ? 1 : -1);
const isSameColor = (c1, c2) => c1 !== undefined && c2 !== undefined
                                && c1.category === c2.category && c1.name === c2.name && c1.code === c2.code;

const mapStateToProps = state => ({
    colorList: state.searchBar.colors,
    baseColor: state.board.baseColor,
});

const mapDispatchToProps = dispatch => ({
    fetchColors: () => dispatch(searchBar.fetchColors()),
    setBaseColor: color => dispatch(board.setBaseColor(color)),
});

export default connect(mapStateToProps, mapDispatchToProps)(SideContent);
