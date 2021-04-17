using System;
using System.Collections.Generic;
using BenchmarkDotNet;
using BenchmarkDotNet.Attributes;

namespace IEnumerableBench
{
    [MemoryDiagnoser]
    public class Benchmarks
    {
        [Params(1000, 1000000, 10000000)]
        public int Size { get; set; }

        private readonly IEnumerable<int> _sequence;
        private readonly IList<int> _list;

        public Benchmarks()
        {
            _sequence = PrepareRandomEnumerable(Size);
            _list = PrepareRandomList(Size);
        }

        [Benchmark]
        public void Sum()
        {
            var sum = 0;

            foreach (var e in _list)
            {
                sum += e;
            }
        }
        
        [Benchmark]
        public void SumByIndexes()
        {
            var sum = 0;

            for (var id = 0; id < _list.Count; id++)
            {
                sum += _list[id];
            }
        }

        [Benchmark]
        public void SumEnumerable()
        {
            var sum = 0;

            foreach (var e in _sequence)
            {
                sum += e;
            }
        }

        private static List<int> PrepareRandomList(int size)
        {
            var list = new List<int>(size);
            var rnd = new Random();

            for (var i = 0; i < size; i++)
            {
                list.Add(rnd.Next(100));
            }
            
            return list;
        }

        private static IEnumerable<int> PrepareRandomEnumerable(int size)
        {
            var rnd = new Random();

            for (var i = 0; i < size; i++)
            {
                yield return rnd.Next(100);
            }
        }
    }
}