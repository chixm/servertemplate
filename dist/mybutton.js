"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var react_1 = __importDefault(require("react"));
var react_dom_1 = __importDefault(require("react-dom"));
var MyButton = /** @class */ (function (_super) {
    __extends(MyButton, _super);
    function MyButton(props) {
        var _this = _super.call(this, props) || this;
        _this.onClickAction = function () {
            alert("clicked! " + _this.state.pressedCount + " times.");
            _this.setState({ pressedCount: _this.state.pressedCount + 1 });
        };
        _this.render = function () {
            return (react_1.default.createElement("div", null,
                react_1.default.createElement("button", { onClick: _this.onClickAction }, _this.props.name)));
        };
        _this.state = { pressedCount: 0 };
        return _this;
    }
    return MyButton;
}(react_1.default.Component));
react_dom_1.default.render(react_1.default.createElement(MyButton, { name: "Count Up" }), document.getElementById('original-button'));
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibXlidXR0b24uanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvdmlldy9teWJ1dHRvbi50c3giXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsZ0RBQTBCO0FBQzFCLHdEQUFpQztBQVdqQztJQUF1Qiw0QkFBNkM7SUFDaEUsa0JBQVksS0FBb0I7UUFBaEMsWUFDSSxrQkFBTSxLQUFLLENBQUMsU0FFZjtRQUVELG1CQUFhLEdBQWU7WUFDeEIsS0FBSyxDQUFDLGNBQVksS0FBSSxDQUFDLEtBQUssQ0FBQyxZQUFZLFlBQVMsQ0FBQyxDQUFDO1lBQ3BELEtBQUksQ0FBQyxRQUFRLENBQUMsRUFBQyxZQUFZLEVBQUUsS0FBSSxDQUFDLEtBQUssQ0FBQyxZQUFZLEdBQUUsQ0FBQyxFQUFFLENBQUMsQ0FBQztRQUMvRCxDQUFDLENBQUE7UUFFRCxZQUFNLEdBQUc7WUFDTCxPQUFPLENBQ0g7Z0JBQ0ksMENBQVEsT0FBTyxFQUFFLEtBQUksQ0FBQyxhQUFhLElBQUcsS0FBSSxDQUFDLEtBQUssQ0FBQyxJQUFJLENBQVUsQ0FDN0QsQ0FDVCxDQUFBO1FBQ0wsQ0FBQyxDQUFBO1FBZEcsS0FBSSxDQUFDLEtBQUssR0FBRyxFQUFDLFlBQVksRUFBQyxDQUFDLEVBQUMsQ0FBQzs7SUFDbEMsQ0FBQztJQWNMLGVBQUM7QUFBRCxDQUFDLEFBbEJELENBQXVCLGVBQUssQ0FBQyxTQUFTLEdBa0JyQztBQUVELG1CQUFRLENBQUMsTUFBTSxDQUFDLDhCQUFDLFFBQVEsSUFBQyxJQUFJLEVBQUUsVUFBVSxHQUFhLEVBQUUsUUFBUSxDQUFDLGNBQWMsQ0FBQyxpQkFBaUIsQ0FBQyxDQUFDLENBQUMifQ==