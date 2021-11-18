import { Select } from "antd";
import { SelectValue } from "antd/lib/select";
import * as React from "react";
const Option = Select.Option;

interface ArrSelectProps {
	className?: string;
	style?: object;
	placeholder?: string;
	names: any[];
	onChange: (value: SelectValue) => void;
	value: SelectValue;
	disabled?: boolean;
	notHasAll?: boolean;
}

export default class ArrSelect extends React.Component<ArrSelectProps, {}> {
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
				allowClear
				defaultValue={notHasAll ? value : "全部"}
				onChange={onChange}
				style={style}
				disabled={disabled}
				placeholder={placeholder}>
				{!notHasAll && (
					<Option key="全部" value="全部">
						全部
					</Option>
				)}
				{names &&
					names.length > 0 &&
					names.map(item => {
						if (item.value && String(item.value)) {
							return (
								<Option key={item.value} value={item.value}>
									{item.label}
								</Option>
							);
						}
						return (
							<Option key={item} value={item}>
								{item}
							</Option>
						);
					})}
			</Select>
		);
	}
}
