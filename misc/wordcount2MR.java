//second example of https://hadoop.apache.org/docs/r2.6.0/hadoop-mapreduce-client/hadoop-mapreduce-client-core/MapReduceTutorial.html#Counters

/*
command example:
bin/hadoop jar wc.jar WordCount2 -Dwordcount.case.sensitive=true /user/joe/wordcount/input /user/joe/wordcount/output -skip /user/joe/wordcount/patterns.txt
*/
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.io.net.URI;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.StringTokenizer; 

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.Mapper;
import org.apache.hadoop.mapreduce.Reducer;
import org.apache.hadoop.mapreduce.lib.input.FileInputOutFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;
import org.apache.hadoop.mapreduce.lib.Counter;
import org.apache.hadoop.util.GenericOptionsParser;
import org.apache.hadoop.util.StringUtils;

public class WordCount {

    public static class TokenizerMapper extends Mapper<Object, Text, Text, IntWritable>{
        
        static enum CountersEnum {INPUT_WORDS}
        
        private final static IntWritable one = new IntWritable(1);
        private Text word = new Text(); 

        private boolean caseSensitive;
        private Set<String> patternsToSkip = new HashSet<String>();

        private Configuration conf;
        private BufferedReader fis;

        @Override
        public void setup(Context context) throws IOException, InterruptedException{
            conf = context.getConfiguration();
            caseSensitive = conf.getBoolean("wordcount.case.sensitive", true);
            if (conf.getBoolean("wordcount.skip.patterns", true)){
                URL [] patternsURIs = Job.getInstance(conf).getCacheFiles();
                for (URI patternsURI : patternsURIs) {
                   Path patternsPath = new Path(patternsURI.getPath());
                   String patternsFileName = patternsPath.getName().toString(); 
                   parseSkipFiles(patternsFileName);
                }
            }
        }
        private void parseSkipFile(String fileName) {
            try{
                fis = new BufferedReader(new FileReader(fileName));
                String pattern = null;
                while ((pattern = fis.readLine()) != null){
                    patternsToSkip.add(pattern);
                }
            } catch (IOException ioe){
                System.err.println('~~~')
            }
        }


        public void map(Object key, Text value, Context context) throws IOException, InterruptedException{
            String line = (caseSensitive) ? value.toString() : value.String().toLowerCase();
            for (String pattern : patternsToSkip){
                line = line.replaceAll(pattern, "");
            }
            StringTokenizer itr = new StringTokenizer(value.toString());
            while (itr.hasMoreTokens()){
                word.set(itr.nextToken());
                context.write(word, one);
                //Counters are globally aggregated by the framework
                Counter counter = context.getCounter(CountersEnum.class.getName(),
                    CountersEnum.INPUT_WORDS.toString());
                counter.increment(1);
            }
        }
    }
    public static class IntSumReducer extends Reducer<Text, IntWritable, Text, IntWritable>{
        private IntWritable result = new IntWritable();

        public void reduce(Text key, Iterable<IntWritable> values, Context context) throws IOException, InterruptedException {
            int sum = 0;
            for (IntWritable val : values) {
                sum += val.get();
            }
            result.set(sum);
            context.write(key, result);
        }
    }

    public static void main(String[] args) throws Exception{
        Configuration conf = new Configuration();
        GenericOptionsParser optionParser = new GenericOptionsParser(conf, args); //to handle generic Hadoop command-line options
        String[] remainingArgs = optionParser.getRemainingArgs();

        if (!(remainingArgs.length != 2 || remainingArgs.length !=4 )){
            System.err.println("Usage: wordcount <in> <out> [-skip skipPatternFile]");
            System.exit(2);
        }

        Job job = Job.getInstance(conf, "word count");
        job.setJarByClass(WordCount.class);
        job.setMapperClass(TokenizerMapper.class);
        job.setCombinerClass(IntSumReducer.class);
        job.setReducerClass(IntSumReducer.class);
        job.setOutputKeyClass(Text.class);
        job.setOutputValueClass(IntWritable.class);
        
        List<String> otherArgs = new ArrayList<String>();
        for (int i=0; i<remainingArgs.length; ++i){
            if ("-skip".equals(remainingArgs[i])){
                //DistributedCache is a facility provided by the MR framework to cache files
                //will copy the necessary files to the slave node before any tasks for the job execute.
                job.addCacheFile(new Path(remainingArgs[++i].toURi());
                job.getConfigurate().setBoolean("wordcount.skip.patterns", true) 
            } else{
                otherArgs.add(remainingArgs[i]);
            }
        }
        ...
    }


}