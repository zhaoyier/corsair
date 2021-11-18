import { Component } from "react";

interface DynamicProps {
	path: string;
}
export default class Dynamic extends Component<DynamicProps, { module: any }> {
	constructor(props) {
		super(props);
		this.state = { module: null };
	}
	componentDidMount() {
		const { path } = this.props;
		import(`${path}`).then(module => this.setState({ module: module.default }));
	}
	render() {
		const { module: Module } = this.state;
		return <Module />;
	}
}
