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
        _this.state = { pressedCount: 0 };
        return _this;
    }
    MyButton.prototype.onClickAction = function () {
        alert("clicked! " + this.state.pressedCount + " times.");
        this.setState({ pressedCount: this.state.pressedCount + 1 });
    };
    MyButton.prototype.render = function () {
        return (react_1.default.createElement("div", null,
            react_1.default.createElement("button", { onClick: this.onClickAction }, this.props.name)));
    };
    return MyButton;
}(react_1.default.Component));
react_dom_1.default.render(react_1.default.createElement(MyButton, { name: "Count Up" }), document.getElementById('original-button'));
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibXlidXR0b24uanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyIuLi9zcmMvdmlldy9teWJ1dHRvbi50c3giXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsZ0RBQTBCO0FBQzFCLHdEQUFpQztBQVdqQztJQUF1Qiw0QkFBNkM7SUFDaEUsa0JBQVksS0FBb0I7UUFBaEMsWUFDSSxrQkFBTSxLQUFLLENBQUMsU0FFZjtRQURHLEtBQUksQ0FBQyxLQUFLLEdBQUcsRUFBQyxZQUFZLEVBQUMsQ0FBQyxFQUFDLENBQUM7O0lBQ2xDLENBQUM7SUFFRCxnQ0FBYSxHQUFiO1FBQ0ksS0FBSyxDQUFDLGNBQVksSUFBSSxDQUFDLEtBQUssQ0FBQyxZQUFZLFlBQVMsQ0FBQyxDQUFDO1FBQ3BELElBQUksQ0FBQyxRQUFRLENBQUMsRUFBQyxZQUFZLEVBQUUsSUFBSSxDQUFDLEtBQUssQ0FBQyxZQUFZLEdBQUUsQ0FBQyxFQUFFLENBQUMsQ0FBQztJQUMvRCxDQUFDO0lBRUQseUJBQU0sR0FBTjtRQUNJLE9BQU8sQ0FDSDtZQUNJLDBDQUFRLE9BQU8sRUFBRSxJQUFJLENBQUMsYUFBYSxJQUFHLElBQUksQ0FBQyxLQUFLLENBQUMsSUFBSSxDQUFVLENBQzdELENBQ1QsQ0FBQTtJQUNMLENBQUM7SUFDTCxlQUFDO0FBQUQsQ0FBQyxBQWxCRCxDQUF1QixlQUFLLENBQUMsU0FBUyxHQWtCckM7QUFFRCxtQkFBUSxDQUFDLE1BQU0sQ0FBQyw4QkFBQyxRQUFRLElBQUMsSUFBSSxFQUFFLFVBQVUsR0FBYSxFQUFFLFFBQVEsQ0FBQyxjQUFjLENBQUMsaUJBQWlCLENBQUMsQ0FBQyxDQUFDIn0=