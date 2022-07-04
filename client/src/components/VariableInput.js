import { Switch, InputNumber, Input } from 'antd';

export const VariableType = {
    Bool: "bool",
    Int: "int",
    Int8: "int8",
    Int16: "int16",
    Int32: "int32",
    Int64: "int64",
    Uint: "uint",
    Uint8: "uint8",
    Uint16: "uint16",
    Uint32: "uint32",
    Uint64: "uint64",
    Float32: "float32",
    Float64: "float64",
    String: "string",
}

function getIntRangeFromBit(bit, sign = true) {
    if (!sign) {
        const valueRange = 2 ** bit
        return [valueRange, 0]
    } else {
        const valueRange = 2 ** bit / 2
        return [valueRange - 1, -1 * valueRange]
    }
}

function getIntRangeFromType(type) {
    switch (type) {
        case VariableType.Int8:
            return getIntRangeFromBit(8)
        case VariableType.Int16:
            return getIntRangeFromBit(16)
        case VariableType.Int32:
            return getIntRangeFromBit(32)
        case VariableType.Int64:
            return getIntRangeFromBit(64)
        case VariableType.Uint:
            return getIntRangeFromBit(32)
        case VariableType.Uint8:
            return getIntRangeFromBit(8, false)
        case VariableType.Uint16:
            return getIntRangeFromBit(16, false)
        case VariableType.Uint32:
            return getIntRangeFromBit(32, false)
        case VariableType.Uint64:
            return getIntRangeFromBit(64, false)
        default:
            return [0, 0]
    }
}

function IntInput({ defaultValue = 0, onChange, type }) {
    const [max, min] = getIntRangeFromType(type)
    return <InputNumber
        style={{ width: "100%" }}
        defaultValue={defaultValue}
        min={String(min)}
        max={String(max)}
        step="1"
        onChange={onChange}
    />
}

function getFloatRangeFromType(type) {
    switch (type) {
        case VariableType.Float32:
        case VariableType.Float64:
            return [3.4E+38, 1.2E-38]
        default:
            return [0, 0]
    }
}

function FloatInput({ defaultValue = 0, onChange, type }) {
    const [max, min] = getFloatRangeFromType(type)
    return <InputNumber
        style={{ width: "100%" }}
        defaultValue={defaultValue}
        min={min}
        max={max}
        onChange={onChange}
        precision={100}
    />
}

function StringInput({ defaultValue = "", onChange }) {
    return <Input style={{ width: "100%" }} defaultValue={defaultValue} onChange={({ target }) => onChange(target.value)}></Input>
}

function BoolInput({ defaultValue = false, onChange }) {
    return <Switch defaultChecked={defaultValue} onChange={({ target }) => onChange(target.checked)} />
}

function VariableInput({ defaultValue, onChange, type }) {
    switch (type) {
        case VariableType.Int:
        case VariableType.Int8:
        case VariableType.Int16:
        case VariableType.Int32:
        case VariableType.Int64:
        case VariableType.Uint:
        case VariableType.Uint8:
        case VariableType.Uint16:
        case VariableType.Uint32:
        case VariableType.Uint64:
            return <IntInput defaultValue={defaultValue} onChange={onChange} type={type} />
        case VariableType.Float32:
        case VariableType.Float64:
            return <FloatInput defaultValue={defaultValue} onChange={onChange} type={type} />
        case VariableType.String:
            return <StringInput defaultValue={defaultValue} onChange={onChange} />
        case VariableType.Bool:
            return <BoolInput defaultValue={defaultValue} onChange={onChange} />
        default:
            return null
    }
}

export default VariableInput 