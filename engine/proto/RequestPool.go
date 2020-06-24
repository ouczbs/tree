package proto

type (
	TSize = uint32
)
type UStackPool struct {
	size TSize

	head TSize
	tail TSize

	pool []IPoolObject
}
func (stack * UStackPool) Len()TSize{
	return (stack.head - stack.tail + stack.size) % stack.size
}
func (stack * UStackPool) Size()TSize{
	return stack.size
}
func (stack * UStackPool) MakeAdd(size TSize){
	stack.head = stack.size
	stack.tail = 0
	stack.size += size
	pool := make([]IPoolObject , size)
	stack.pool = append(stack.pool , pool...)
}

func (stack * UStackPool) Push(object IPoolObject){
	object.Release()
	head := stack.head
	stack.pool[head] = object
	stack.head = ( head + 1 ) % stack.size
	if stack.head == stack.tail { //栈满
		stack.MakeAdd(stack.size)
	}
}

func (stack * UStackPool) Pop() IPoolObject{
	tail := stack.tail
	if tail == stack.head { //栈空
		return nil
	}
	stack.tail = (tail + stack.size - 1) % stack.size
	return stack.pool[tail]
}

type URequest struct{
	// Response
	ProtoMessage IReflectMessage
	Code         TCode
	Next		 TFlag
	// Request
	MessageType  TMessageType
	Cmd 		 TCmd
	request      TCallId
}
func (request * URequest) Release(){
	request.Code = CodeOk
	request.Cmd = 0
	request.MessageType = 0
	request.Next = true
}
type URequestPool struct {
	*UStackPool
}

func (stack * URequestPool) CreateRequest() * URequest{
	request := new(URequest)
	return request
}
func (stack * URequestPool) Pop() * URequest{
	var request * URequest
	object := stack.UStackPool.Pop()
	if object != nil  {
		request = object.( * URequest)
	}else {
		request = stack.CreateRequest()
	}
	return request
}
func NewRequestPool(size TSize)* URequestPool{
	stack := &URequestPool{&UStackPool{size: size}}
	stack.pool = make([]IPoolObject, size)
	return stack
}