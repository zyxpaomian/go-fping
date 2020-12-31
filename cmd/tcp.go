package cmd

import (
	"fmt"
	"net"
	"bufio"
	"io"	
	"os"
	"github.com/spf13/cobra"
	"sync"
	"time"
)

// tcpCmd represents the tcp command
var tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "tcp fping network",
	Long: `tcp fping network`,
	Run: func(cmd *cobra.Command, args []string) {
		var ipList []string

		// 探测源，只能选1
		if len(file) != 0 && (len(subnet) != 0 || len(singleip) != 0) || len(subnet) != 0 && (len(file) != 0 || len(singleip) != 0) || len(singleip) != 0 && (len(file) != 0 || len(subnet) != 0){
			fmt.Println("请检查输入参数, -f, -g, -i 不能共存")
			cmd.Help()
			return 
		}
		
		// 检查输出文件
		if len(output) != 0 {
			if _, err := os.Stat(output); err != nil {
				f, err := os.Create(output)
				defer f.Close()
				if err != nil {
					fmt.Println("结果输出文件无法创建")
        			return					
				}
			}
			outputFile, err := os.Open(output)
    		if err != nil {
				fmt.Println("结果输出文件无法打开")
        		return
    		}
			outputFile.Close()
		}
					
		// 探测文件IP得情况
		if len(file) != 0 {
			fi, err := os.Open(file)
    		if err != nil {
				fmt.Println("IP列表文件无法打开")
        		return
    		}
    		defer fi.Close()

    		br := bufio.NewReader(fi)
    		for {
        		a, _, c := br.ReadLine()
        		if c == io.EOF {
           			break
        		}
			ipList = append(ipList, string(a))
    		}
		}

		// 探测网段得情况下
		if len(subnet) != 0 {
			ips, err := subNetGet(subnet)
			if err != nil {
				panic(err)
			}
			ipList = ips
		}

		// 单个IP得情况
		if len(singleip) != 0 {
			ipList = append(ipList, singleip)
		}


		// 实际的Tcp操作
		var reachableIps []string
		var unreachableIps []string

		var lock sync.Mutex
		var wg sync.WaitGroup

    	buckets := make(chan bool, routinepool)
		for _, ip := range(ipList) {
			buckets <- true
			wg.Add(1)
			go func(ip string) {
				result, err := tcpAlive(ip)
  				lock.Lock()
  				defer lock.Unlock()				
				if err != nil {
					unreachableIps = append(unreachableIps, ip)
				} else {
					if result == true {
						reachableIps = append(reachableIps, ip)
					} else {
						unreachableIps = append(unreachableIps, ip)
					}
				}
				<- buckets
				wg.Done()	
			}(ip)
		}
		wg.Wait()

		if len(output) != 0 {	
			outputFile, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, os.ModePerm)
    		if err != nil {
				fmt.Println("结果输出文件无法打开")
        		return
    		}
    		defer outputFile.Close()

			for _, rip := range(reachableIps) {
				outputFile.WriteString(rip + ": ok\n")
			}
			for _, urip := range(unreachableIps) {
				outputFile.WriteString(urip + ": failed\n")
			}
			fmt.Printf("结果输出到文件: %s \n",output)
		} else {
			fmt.Printf("Reachable IP: %v\n", reachableIps)
			fmt.Printf("UnReachable IP: %v\n", unreachableIps)
		}
	},
}


func tcpAlive(ip string) (bool, error) {
	tcpTimeOut := time.Duration(timeout)
    conn, err := net.DialTimeout("tcp", ip, time.Duration(tcpTimeOut * time.Millisecond))
    if err != nil {
		return false, err
    }
    conn.Close()
	return true, nil
}

func init() {
	rootCmd.AddCommand(tcpCmd)
}
