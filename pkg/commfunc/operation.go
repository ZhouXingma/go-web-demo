package commfunc

// 模拟三元操作
// 参数：
// bool：判断是否成功
// trueValue：成功的值
// falseValue：失败的值
// 返回：
// T: 根据bool计算得到的结果
func BoolHandle[T any](bool bool, trueValue T, falseValue T) T {
	if bool {
		return trueValue
	}
	return falseValue
}
