// ---------
// CPU CACHE
// ---------

/*

 Cores DO NOT access main memory directly but their local caches.
 (Core không truy cập trực tiếp vào main memory [bộ nhớ chính] mà là local caches [bộ nhớ đệm cục bộ] của chúng)

 What store in caches are data and instruction.
 (Những gì lưu trữ trong cache là data và instruction [lệnh chỉ thị])

 Cache speed from fastest to slowest: L1 -> L2 -> L3 -> main memory.
 (Tốc đọ cache từ nhanh đến chậm: L1 -> L2 -> L3 -> main memory.)

 Scott Meyers: "If performance matter then the total memory you have is the total amount of
 caches" -> access to main memory is incredibly slow; practically speaking it might not even be there.
 (Scott Meyers: 'Nếu hiệu suất quang trọng thì tổng bộ nhớ bạn có là tổng dung lượng bộ nhớ cache [đệm] '
  -> truy cập từ main memory cực kỳ chậm; thực tế mà nói thì thậm chí có thể không có ở đó)

 How do we write code that can be sympathetic with the caching system to make sure that
 we don't have a cache miss or at least, we minimize cache misses to our fullest potential?
 (chúng ta viết code thế nào để tối ưu và giảm thiểu cache misses.)

 Processor has a Prefetcher. It predicts what data is needed ahead of time.
 (Processor có trình đặt trước. Nó có thể dự đoán data nào cần trước.)

 There are different granularity depending on where we are on the machine.
(Có nhiều độ chi tiết khác nhau tùy thuộc vào vị ví của chúng ta trên máy. )

 Our programming model uses a byte. We can read and write to a byte at a time. However, from the
 caching system POV, our granularity is not 1 byte. It is 64 bytes, called a cache line. All
 memory is junked up in this 64 bytes cache line.
 (Mô hình lập trình của chúng tối sử dụng 1 byte. Chúng ta có thể đọc và viết một byte tại một thời điểm. Tuy nhiên,
 từ cache system POV, không phải là 1 byte. Nó là 64 bytes, gọi là cache line. Tất cả memory được tăng lên trong
 64 bytes cache line này.)


 Since the caching mechanism is complex, Prefetcher tries to hide all the latency from us.
 (Cơ chế caching rất phức tạp nên Prefetcher cố gắng ẩn tất cả độ trễ.)

 It has to be able to pick up on predictable access pattern to data.
 (Nó có khả năng nhận diện được access pattern đến data)
 -> We need to write code that creates predictable access pattern to data
 (Chúng ta cần viết code tạo một access pattern đến data có thể dự đoán được.)

 One easy way is to create a contiguous allocation of memory and to iterate over them.
 (Một cách dễ dàng tạo phân bổ bộ nhớ liền kề và lặp lại chúng.)

 The array data structure gives us ability to do so.
 (cấu trúc Array có thể làm như vậy.)

 From the hardware perspective, array is the most important data structure.
 (Từ góc độ phần cứng, Array là cấu trúc dữ liệu quan trọng bậc nhất.)

 From Go perspective, slice is. Array is the backing data structure for slice (like Vector in C++).
 (Tù Go thì nó là slice. Array là cấu trúc dữ liệu bổ trợ cho slice [như Vectỏ trong C++].)

 Once we allocate an array, whatever it size, every element is equal distant from other element.
 (Khi chúng ta phân bổ một array, bất kể kích thước nào, mỗi element đều tương đương với element khác.)

  As we iterate over that array, we begin to walk cache line by cache line. As the Prefetcher see
  that access pattern, it can pick it up and hide all the latency from us.
  (Như chúng ta lặp mảng, chúng ta bắt đầu đi qua từng cache line. Prefetcher nhìn thấy access pattern đó,
   nó có thể nhận ra và ẩn tất cả độ trễ với chúng ta.)

 For example, we have a big nxn matrix. We do LinkedList Traverse, Column Traverse, and Row Traverse
 and benchmark against them.
 (Ví dụ, chúng ta có một ma trận lớn nxn. Chúng thực hiện LinkedList Traverse, Column Traverse, và row Traverse
 và cũng benchmark chúng.)

 Unsurprisingly, Row Traverse has the best performance. It walk through the matrix cache line
 by cache line and create a predictable access pattern.
 (Chẳng có gì ngạc nhiên, Row Traverse có hiệu suất tốt nhất. Nó đi qua matrix cache line theo cache line
 và tạo một predictable access pattern.)

 Column Traverse does not walk through the matrix cache line by cache line. It looks like random
 access memory pattern. That is why is the slowest among those.
 (Clonumn Traverse không đi qua matrix cache line theo cache line. Nó giống như random access memory pattern.
 Đó là lý do tại sao nó chậm nhất.)

 However, that doesn't explain why the LinkedList Traverse's performance is in the middle. We
 just think that it might perform as poorly as the Column Traverse.
 -> This leads us to another cache: TLB - Translation lookaside buffer. Its job is to maintain
 operating system page and offset to where physical memory is.
 (LinkedList Traverse dẫn chúng tôi đến một cache khác: TLB - Translation lookaside buffer. Công việc của nó
 là duy trì page của hệ điều hành và offet (bù đắp) vào bộ nhớ vật lý.)

// ----------------------------
// Translation lookaside buffer
// ----------------------------

 Back to the different granularity, the caching system moves data in and out the hardware at 64
 bytes at a time. However, the operating system manages memory by paging its 4K (traditional page
 size for an operating system).
 (Trở lại với mức độ chi tiết khác, cache system di chuyển data vào và ra phần cứng 64 bytes tại
 một thời điểm. Tuy nhiên hệ điều hành quản lý bộ nhớ bởi paging là 4k [kích thước page truyền thống
 cho hệ điều hành].)

 TLB: For every page that we are managing, let's take our virtual memory addresses that we use
 (softwares run virtual addresses, its sandbox, that is how we use/share physical memory)
 and map it to the right page and offset for that physical memory.


 A miss on the TLB can be worse than just the cache miss alone.
 The LinkedList is somewhere in between is because the chance of multiple nodes being on the same
 page is probably pretty good. Even though we can get cache misses because cache lines aren't
 necessary in the distance that is predictable, we probably don't have so many TLB cache misses.
 In the Column Traverse, not only we have cache misses, we probably have a TLB cache miss on
 every access as well.

 Data-oriented design matters.
 It is not enough to write the most efficient algorithm, how we access our data can have much
 more lasting effect on the performance than the algorithm itself.

*/
package main

import "fmt"

func main() {
	// -----------------------
	// Declare and initialize
	// -----------------------

	// Declare an array of five strings that is initialized to its zero value.
	// Recap: a string is a 2 word data structure: a pointer and a length
	// Since this array is set to its zero value, every string in this array is also set to its
	// zero value, which means that each string has the first word pointed to nil and
	// second word is 0.
	//  -----------------------------
	// | nil | nil | nil | nil | nil |
	//  -----------------------------
	// |  0  |  0  |  0  |  0  |  0  |
	//  -----------------------------
	var strings [5]string

	// At index 0, a string now has a pointer to a backing array of bytes (characters in string)
	// and its length is 5.

	// -----------------
	// What is the cost?
	// -----------------

	// The cost of this assignment is the cost of copying 2 bytes.
	// We have two string values that have pointers to the same backing array of bytes.
	// Therefore, the cost of this assignment is just 2 words.

	//  -----         -------------------
	// |  *  |  ---> | A | p | p | l | e | (1)
	//  -----         -------------------
	// |  5  |                  A
	//  -----                   |
	//                          |
	//                          |
	//     ---------------------
	//    |
	//  -----------------------------
	// |  *  | nil | nil | nil | nil |
	//  -----------------------------
	// |  5  |  0  |  0  |  0  |  0  |
	//  -----------------------------
	strings[0] = "Apple"
	strings[1] = "Orange"
	strings[2] = "Banana"
	strings[3] = "Grape"
	strings[4] = "Plum"

	// ---------------------------------
	// Iterate over the array of strings
	// ---------------------------------

	// Using range, not only we can get the index but also a copy of the value in the array.
	// fruit is now a string value; its scope is within the for statement.
	// In the first iteration, we have the word "Apple". It is a string that has the first word
	// also points to (1) and the second word is 5.
	// So we now have 3 different string value all sharing the same backing array.

	// What are we passing to the Println function?
	// We are using value semantic here. We are not sharing our string value. Println is getting
	// its own copy, its own string value. It means when we get to the Println call, there are now
	// 4 string values all sharing the same backing array.

	// We don't want to take an address of a string.
	// We know the size of a string ahead of time.
	// -> it has the ability to be on the stack
	// -> not creating allocation
	// -> not causing pressure on the GC
	// -> the string has been designed to leverage value mechanic, to stay on the stack, out of the
	// way of creating garbage.
	// -> the only thing that has to be on the heap, if anything is the backing array, which is the
	// one thing that being shared
	fmt.Printf("\n=> Iterate over array\n")
	for i, fruit := range strings {
		fmt.Println(i, fruit)
	}

	// Declare an array of 4 integers that is initialized with some values using literal syntax.
	numbers := [4]int{10, 20, 30, 40}

	// Iterate over the array of numbers using traditional style.
	fmt.Printf("\n=> Iterate over array using traditional style\n")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	// ---------------------
	// Different type arrays
	// ---------------------

	// Declare an array of 5 integers that is initialized to its zero value.
	var five [5]int

	// Declare an array of 4 integers that is initialized with some values.
	four := [4]int{10, 20, 30, 40}

	fmt.Printf("\n=> Different type arrays\n")
	fmt.Println(five)
	fmt.Println(four)

	// When we try to assign four to five like so five = four, the compiler says that
	// "cannot use four (type [4]int) as type [5]int in assignment"
	// This cannot happen because they have different types (size and representation).
	// The size of an array makes up its type name: [4]int vs [5]int. Just like what we've seen
	// with pointer. The * in *int is not an operator but part of the type name.

	// Unsurprisingly, all array has known size at compiled time.

	// -----------------------------
	// Contiguous memory allocations
	// -----------------------------

	// Declare an array of 6 strings initialized with values.
	six := [6]string{"Annie", "Betty", "Charley", "Doug", "Edward", "Hoanh"}

	// Iterate over the array displaying the value and address of each element.
	// By looking at the output of this Printf function, we can see that this array is truly a
	// contiguous block of memory. We know a string is 2 word and depending on computer
	// architecture, it will have x byte. The distance between two consecutive IndexAddr is exactly
	// x byte.
	// v is its own variable on the stack and it has the same address every single time.
	fmt.Printf("\n=> Contiguous memory allocations\n")
	for i, v := range six {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &six[i])
	}
}
