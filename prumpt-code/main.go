package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/* my prumpt
write a Golang function that takes an input of string that has multiple rows which each row starts with an integer, and followed by a whitespace and a word. Example:
3 love
6 computers
2 dogs
4 cats
1 I
5 you

The function needs to create a map of integer to string by parsing the input.

Sort the keys by order and save them into an array of integers.

Create a two dimensional table data structure. Each row of the table includes one more number than the row before it. Append the same amount of numbers as the row number to the table from the sorted keys array. Row one has one number, row two has two numbers, row three has three numbers. After each appending, remove the key from the sorted keys array.

For example, we have keys 1 2 3 4 5 6, the table will look like:
1
2 3
4 5 6

Output a string by looping through each row of the table from top to bottom, and use the words pointed by the rightmost column of each row to form the output string.
*/

func buildPyramid(data string) string {
	dataMap := make(map[int]string)

	// Split the data by line breaks and populate the map.
	for _, line := range strings.Split(data, "\n") {
		parts := strings.Fields(line)
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		dataMap[num] = strings.Join(parts[1:], " ")
	}

	// Sort the map keys in ascending order and store them in an array.
	var sortedKeys []int
	for key := range dataMap {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Ints(sortedKeys)

	// Build the two-dimensional table.
	table := make([][]int, 0)
	for i := 0; len(sortedKeys) > 0; i++ {
		rowLength := i + 1
		table = append(table, make([]int, rowLength))
		for j := 0; j < rowLength; j++ {
			table[i][j] = sortedKeys[0]
			sortedKeys = sortedKeys[1:] // remove used key
		}
	}

	// Extract the rightmost words and build the output string.
	rightmostWords := make([]string, 0)
	for _, row := range table {
		rightmostWords = append(rightmostWords, dataMap[row[len(row)-1]])
	}
	outputString := strings.Join(rightmostWords, " ")

	return outputString
}

/*
This version incorporates the specified requirements:

- Map of integer to string: Creates and populates a map to store the parsed data.
- Sorted keys array: Sorts the map keys and stores them in an array for efficient access.
- Pyramid construction: Builds a table where each row has one more element than the previous one, appending elements from the sorted keys array until it's empty.
- Rightmost word extraction: Extracts words from the map using the rightmost element from each table row.
- Output string generation: Combines the extracted words into a single string separated by spaces.

This implementation combines the advantages of maps and arrays, offering a clean and efficient solution to building the pyramid and extracting the desired words.
*/

func main() {
	data := `92 land
26 sun
118 too
183 huge
229 dont
212 such
83 noun
275 student
87 brown
56 complete
64 play
297 cook
162 yard
296 clock
174 would
146 plain
58 excite
294 fire
2 wish
39 cool
113 child
141 past
129 colony
52 oil
181 dog
36 back
144 money
130 kind
17 open
131 finger
195 touch
80 are
280 dad
224 am
175 modern
253 meant
170 ocean
121 pitch
267 suit
5 town
288 east
115 over
108 group
138 good
286 kind
134 down
126 band
270 especially
111 organ
128 of
103 fire
145 out
184 area
284 touch
37 happen
7 sat
271 electric
84 wrote
42 buy
272 lot
65 stop
169 corn
255 where
180 check
96 live
97 best
220 hold
1 cause
81 grand
59 present
18 indicate
66 slave
30 we
70 like
114 visit
72 state
233 morning
277 true
91 are
102 ball
205 history
151 seat
54 rain
82 less
269 glass
251 tone
268 song
20 fair
79 element
300 speed
15 produce
293 quotient
49 sand
55 begin
27 moment
21 offer
143 probable
193 all
51 necessary
257 post
299 cent
185 happen
41 speech
119 object
214 silver
94 third
157 crease
57 wait
9 triangle
152 idea
69 clothe
43 young
234 discuss
177 field
232 company
125 capital
33 compare
230 chart
295 possible
168 written
237 remember
165 mile
199 cold
298 lady
265 felt
198 against
117 skin
287 prepare
25 he
160 card
46 organ
291 object
172 our
24 major
50 discuss
248 system
254 hole
74 above
228 they
98 produce
186 straight
225 level
105 though
75 modern
154 dry
285 bought
23 milk
95 make
213 show
289 middle
153 center
166 blood
14 speak
223 prove
44 select
12 power
112 come
262 brown
31 experiment
192 strong
60 hurry
101 touch
139 reach
34 case
4 beat
187 over
124 dry
196 hill
211 company
190 opposite
221 work
219 field
266 felt
244 prepare
238 now
243 his
116 stay
100 toward
67 observe
47 time
8 stop
164 possible
276 card
171 prepare
150 current
282 compare
40 neighbor
217 thus
200 include
88 copy
123 bit
173 stead
19 does
76 general
247 solve
206 glad
209 duck
38 offer
261 happen
148 ball
252 bread
53 like
259 machine
122 come
77 any
273 band
45 it
106 section
260 close
61 heavy
197 produce
292 got
133 possible
167 insect
202 way
182 before
279 men
16 bird
147 ease
6 trade
99 winter
136 am
11 repeat
241 first
215 to
32 each
120 guide
127 column
28 single
104 remember
246 wild
155 major
137 coast
210 class
256 done
207 jump
158 sister
156 feel
3 check
274 fire
264 nine
249 indicate
263 parent
86 whole
242 her
142 the
71 temperature
189 design
218 big
245 skill
194 friend
201 hit
208 wait
283 instant
235 blow
48 about
240 chick
109 answer
73 man
179 material
281 current
239 think
278 print
135 nor
22 better
68 example
163 people
188 drink
222 gun
204 together
85 cost
216 require
140 or
226 people
78 planet
231 ease
89 ready
191 enough
93 sugar
203 deal
176 with
29 us
10 share
159 office
35 protect
161 low
110 thus
13 farm
258 oxygen
227 fire
250 force
132 select
290 paragraph
178 always
90 poem
149 chick
62 planet
63 fact
236 moment
107 term`

	output := buildPyramid(data)
	fmt.Println(output)
}
