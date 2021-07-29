/*
方法	                            寻找消息边界的方式	                    优点	缺点
TCP连接改为短连接，一个请求一个短连接	建立连接释放连接之间的信息即为传输的信息	简单	效率低下
固定长度	                        消息统一满足固定长度，不足补零或者其他	    简单	浪费空间
分隔符	                        分隔符之间	                        空间不浪费，也比较简单	内容本身出现分隔符需要转义,需要扫描全部内容
固定长度字段存内容的长度信息	        先解析固定长度的字段获取长度，然后读取后续内容	精确定位数据，内容不用转义	长度理论上有限制，需提前预支可能的最大长度从而定义长度占用字节
其他方式	例如JSON以{}是否成对出现

*/
package main