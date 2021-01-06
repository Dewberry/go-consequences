package crops

//NASSCropMap is a map of Crops to NASS Crop ID #s
func NASSCropMap() map[string]Crop {
	m := make(map[string]Crop)

	m["1"] = BuildCrop(1, "Corn")
	m["2"] = BuildCrop(2, "Cotton")
	m["3"] = BuildCrop(3, "Rice")
	m["4"] = BuildCrop(4, "Sorghum")
	m["5"] = BuildCrop(5, "Soybeans")
	m["6"] = BuildCrop(6, "Sunflower")
	m["10"] = BuildCrop(10, "Peanuts")
	m["11"] = BuildCrop(11, "Tobacco")
	m["12"] = BuildCrop(12, "Sweet Corn")
	m["13"] = BuildCrop(13, "Pop or Orn Corn")
	m["14"] = BuildCrop(14, "Mint")
	m["21"] = BuildCrop(21, "Barley")
	m["22"] = BuildCrop(22, "Durum Wheat")
	m["23"] = BuildCrop(23, "Spring Wheat")
	m["24"] = BuildCrop(24, "Winter Wheat")
	m["25"] = BuildCrop(25, "Other Small Grains")
	m["26"] = BuildCrop(26, "Dbl Crop WinWht/Soybeans")
	m["27"] = BuildCrop(27, "Rye")
	m["28"] = BuildCrop(28, "Oats")
	m["29"] = BuildCrop(29, "Millet")
	m["30"] = BuildCrop(30, "Speltz")
	m["31"] = BuildCrop(31, "Canola")
	m["32"] = BuildCrop(32, "Flaxseed")
	m["33"] = BuildCrop(33, "Safflower")
	m["34"] = BuildCrop(34, "Rape Seed")
	m["35"] = BuildCrop(35, "Mustard")
	m["36"] = BuildCrop(36, "Alfalfa")
	m["37"] = BuildCrop(37, "Other Hay/Non Alfalfa")
	m["38"] = BuildCrop(38, "Camelina")
	m["39"] = BuildCrop(39, "Buckwheat")
	m["41"] = BuildCrop(41, "Sugarbeets")
	m["42"] = BuildCrop(42, "Dry Beans")
	m["43"] = BuildCrop(43, "Potatoes")
	m["44"] = BuildCrop(44, "Other Crops")
	m["45"] = BuildCrop(45, "Sugarcane")
	m["46"] = BuildCrop(46, "Sweet Potatoes")
	m["47"] = BuildCrop(47, "Misc Vegs & Fruits")
	m["48"] = BuildCrop(48, "Watermelons")
	m["49"] = BuildCrop(49, "Onions")
	m["50"] = BuildCrop(50, "Cucumbers")
	m["51"] = BuildCrop(51, "Chick Peas")
	m["52"] = BuildCrop(52, "Lentils")
	m["53"] = BuildCrop(53, "Peas")
	m["54"] = BuildCrop(54, "Tomatoes")
	m["55"] = BuildCrop(55, "Caneberries")
	m["56"] = BuildCrop(56, "Hops")
	m["57"] = BuildCrop(57, "Herbs")
	m["58"] = BuildCrop(58, "Clover/Wildflowers")
	m["59"] = BuildCrop(59, "Sod/Grass Seed")
	m["60"] = BuildCrop(60, "Switchgrass")
	m["61"] = BuildCrop(61, "Fallow/Idle Cropland")
	m["63"] = BuildCrop(63, "Forest")
	m["64"] = BuildCrop(64, "Shrubland")
	m["65"] = BuildCrop(65, "Barren")
	m["66"] = BuildCrop(66, "Cherries")
	m["67"] = BuildCrop(67, "Peaches")
	m["68"] = BuildCrop(68, "Apples")
	m["69"] = BuildCrop(69, "Grapes")
	m["70"] = BuildCrop(70, "Christmas Trees")
	m["71"] = BuildCrop(71, "Other Tree Crops")
	m["72"] = BuildCrop(72, "Citrus")
	m["74"] = BuildCrop(74, "Pecans")
	m["75"] = BuildCrop(75, "Almonds")
	m["76"] = BuildCrop(76, "Walnuts")
	m["77"] = BuildCrop(77, "Pears")
	m["92"] = BuildCrop(92, "Aquaculture")
	m["152"] = BuildCrop(152, "Shrubland")
	m["204"] = BuildCrop(204, "Pistachios")
	m["205"] = BuildCrop(205, "Triticale")
	m["206"] = BuildCrop(206, "Carrots")
	m["207"] = BuildCrop(207, "Asparagus")
	m["208"] = BuildCrop(208, "Garlic")
	m["209"] = BuildCrop(209, "Cantaloupes")
	m["210"] = BuildCrop(210, "Prunes")
	m["211"] = BuildCrop(211, "Olives")
	m["212"] = BuildCrop(212, "Oranges")
	m["213"] = BuildCrop(213, "Honeydew Melons")
	m["214"] = BuildCrop(214, "Broccoli")
	m["215"] = BuildCrop(215, "Avocados")
	m["216"] = BuildCrop(216, "Peppers")
	m["217"] = BuildCrop(217, "Pomegranates")
	m["218"] = BuildCrop(218, "Nectarines")
	m["219"] = BuildCrop(219, "Greens")
	m["220"] = BuildCrop(220, "Plums")
	m["221"] = BuildCrop(221, "Strawberries")
	m["222"] = BuildCrop(222, "Squash")
	m["223"] = BuildCrop(223, "Apricots")
	m["224"] = BuildCrop(224, "Vetch")
	m["225"] = BuildCrop(225, "Dbl Crop WinWht/Corn")
	m["226"] = BuildCrop(226, "Dbl Crop Oats/Corn")
	m["227"] = BuildCrop(227, "Lettuce")
	m["228"] = BuildCrop(228, "Dbl Crop Triticale/Corn")
	m["229"] = BuildCrop(229, "Pumpkins")
	m["230"] = BuildCrop(230, "Dbl Crop Lettuce/Durum Wheat")
	m["231"] = BuildCrop(231, "Dbl Crop Lettuce/Cantaloupe")
	m["232"] = BuildCrop(232, "Dbl Crop Lettuce/Cotton")
	m["233"] = BuildCrop(233, "Dbl Crop Lettuce/Barley")
	m["234"] = BuildCrop(234, "Dbl Crop Durum Wht/Sorghum")
	m["235"] = BuildCrop(235, "Dbl Crop Barley/Sorghum")
	m["236"] = BuildCrop(236, "Dbl Crop WinWht/Sorghum")
	m["237"] = BuildCrop(237, "Dbl Crop Barley/Corn")
	m["238"] = BuildCrop(238, "Dbl Crop WinWht/Cotton")
	m["239"] = BuildCrop(239, "Dbl Crop Soybeans/Cotton")
	m["240"] = BuildCrop(240, "Dbl Crop Soybeans/Oats")
	m["241"] = BuildCrop(241, "Dbl Crop Corn/Soybeans")
	m["242"] = BuildCrop(242, "Blueberries")
	m["243"] = BuildCrop(243, "Cabbage")
	m["244"] = BuildCrop(244, "Cauliflower")
	m["245"] = BuildCrop(245, "Celery")
	m["246"] = BuildCrop(246, "Radishes")
	m["247"] = BuildCrop(247, "Turnips")
	m["248"] = BuildCrop(248, "Eggplants")
	m["249"] = BuildCrop(249, "Gourds")
	m["250"] = BuildCrop(250, "Cranberries")
	m["254"] = BuildCrop(254, "Dbl Crop Barley/Soybeans")

	return m
}

/*
value,red,green,blue,category,opacity
1	255	211	0	Corn	255
2	255	38	38	Cotton	255
3	0	168	228	Rice	255
4	255	158	11	Sorghum	255
5	38	112	0	Soybeans	255
6	255	255	0	Sunflower	255
10	112	165	0	Peanuts	255
11	0	175	75	Tobacco	255
12	221	165	11	Sweet Corn	255
13	221	165	11	Pop or Orn Corn	255
14	126	211	255	Mint	255
21	226	0	124	Barley	255
22	137	98	84	Durum Wheat	255
23	216	181	107	Spring Wheat	255
24	165	112	0	Winter Wheat	255
25	214	158	188	Other Small Grains	255
26	112	112	0	Dbl Crop WinWht/Soybeans	255
27	172	0	124	Rye	255
28	160	89	137	Oats	255
29	112	0	73	Millet	255
30	214	158	188	Speltz	255
31	209	255	0	Canola	255
32	126	153	255	Flaxseed	255
33	214	214	0	Safflower	255
34	209	255	0	Rape Seed	255
35	0	175	75	Mustard	255
36	255	165	226	Alfalfa	255
37	165	242	140	Other Hay/Non Alfalfa	255
38	0	175	75	Camelina	255
39	214	158	188	Buckwheat	255
41	168	0	228	Sugarbeets	255
42	165	0	0	Dry Beans	255
43	112	38	0	Potatoes	255
44	0	175	75	Other Crops	255
45	177	126	255	Sugarcane	255
46	112	38	0	Sweet Potatoes	255
47	255	102	102	Misc Vegs & Fruits	255
48	255	102	102	Watermelons	255
49	255	204	102	Onions	255
50	255	102	102	Cucumbers	255
51	0	175	75	Chick Peas	255
52	0	221	175	Lentils	255
53	84	255	0	Peas	255
54	242	163	119	Tomatoes	255
55	255	102	102	Caneberries	255
56	0	175	75	Hops	255
57	126	211	255	Herbs	255
58	232	191	255	Clover/Wildflowers	255
59	175	255	221	Sod/Grass Seed	255
60	0	175	75	Switchgrass	255
61	191	191	119	Fallow/Idle Cropland	255
63	147	204	147	Forest	255
64	198	214	158	Shrubland	255
65	204	191	163	Barren	255
66	255	0	255	Cherries	255
67	255	142	170	Peaches	255
68	186	0	79	Apples	255
69	112	68	137	Grapes	255
70	0	119	119	Christmas Trees	255
71	177	154	112	Other Tree Crops	255
72	255	255	126	Citrus	255
74	181	112	91	Pecans	255
75	0	165	130	Almonds	255
76	233	214	175	Walnuts	255
77	177	154	112	Pears	255
92	0	255	255	Aquaculture	255
152	198	214	158	Shrubland	255
204	0	255	140	Pistachios	255
205	214	158	188	Triticale	255
206	255	102	102	Carrots	255
207	255	102	102	Asparagus	255
208	255	102	102	Garlic	255
209	255	102	102	Cantaloupes	255
210	255	142	170	Prunes	255
211	51	73	51	Olives	255
212	228	112	38	Oranges	255
213	255	102	102	Honeydew Melons	255
214	255	102	102	Broccoli	255
215	102	153	76	Avocados	255
216	255	102	102	Peppers	255
217	177	154	112	Pomegranates	255
218	255	142	170	Nectarines	255
219	255	102	102	Greens	255
220	255	142	170	Plums	255
221	255	102	102	Strawberries	255
222	255	102	102	Squash	255
223	255	142	170	Apricots	255
224	0	175	75	Vetch	255
225	255	211	0	Dbl Crop WinWht/Corn	255
226	255	211	0	Dbl Crop Oats/Corn	255
227	255	102	102	Lettuce	255
228	255	210	0	Dbl Crop Triticale/Corn	255
229	255	102	102	Pumpkins	255
230	137	98	84	Dbl Crop Lettuce/Durum Wht	255
231	255	102	102	Dbl Crop Lettuce/Cantaloupe	255
232	255	38	38	Dbl Crop Lettuce/Cotton	255
233	226	0	124	Dbl Crop Lettuce/Barley	255
234	255	158	11	Dbl Crop Durum Wht/Sorghum	255
235	255	158	11	Dbl Crop Barley/Sorghum	255
236	165	112	0	Dbl Crop WinWht/Sorghum	255
237	255	211	0	Dbl Crop Barley/Corn	255
238	165	112	0	Dbl Crop WinWht/Cotton	255
239	38	112	0	Dbl Crop Soybeans/Cotton	255
240	38	112	0	Dbl Crop Soybeans/Oats	255
241	255	211	0	Dbl Crop Corn/Soybeans	255
242	0	0	153	Blueberries	255
243	255	102	102	Cabbage	255
244	255	102	102	Cauliflower	255
245	255	102	102	Celery	255
246	255	102	102	Radishes	255
247	255	102	102	Turnips	255
248	255	102	102	Eggplants	255
249	255	102	102	Gourds	255
250	255	102	102	Cranberries	255
254	38	112	0	Dbl Crop Barley/Soybeans	255
*/
