import { Select } from "antd";
import { SelectValue } from "antd/lib/select";
import * as React from "react";
const Option = Select.Option;

interface ObjectSelectProps {
	className?: string;
	placeholder?: string;
	names: object;
	style?: object;
	onChange: (value: SelectValue) => void;
	value: SelectValue;
	disabled?: boolean;
	notHasAll?: boolean;
}

export default class ObjectSelect extends React.Component<ObjectSelectProps, {}> {
	render() {
		const {
			className,
			names,
			onChange,
			value,
			placeholder,
			style,
			disabled,
			notHasAll
		} = this.props;
		return (
			<Select
				className={className}
				value={value}
				style={style}
				allowClear
				defaultValue={notHasAll ? value : "全部"}
				onChange={onChange}
				disabled={disabled}
				placeholder={placeholder}>
				{!notHasAll && (
					<Option key="全部" value="全部">
						全部
					</Option>
				)}
				{Object.keys(names).map(item => (
					<Option key={item} value={item}>
						{names[item]}
					</Option>
				))}
			</Select>
		);
	}
}
