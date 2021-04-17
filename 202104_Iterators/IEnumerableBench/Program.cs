using System;
using BenchmarkDotNet.Running;

namespace IEnumerableBench
{
    class Program
    {
        static void Main(string[] args)
        {
            var summary = BenchmarkRunner.Run<Benchmarks>();
        }
    }
}
