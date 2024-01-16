package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.R
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.databinding.FragmentInfoChapter1ContentBinding

class InfoChapter1Content : Fragment() {

    companion object {
        fun newInstance() = InfoChapter1Content()
    }

    private lateinit var viewModel: InfoChapter1ContentViewModel
    private var _binding: FragmentInfoChapter1ContentBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
       val infoChapter1ContentViewModel = ViewModelProvider(this)[AccountMainViewModel::class.java]
        _binding = FragmentInfoChapter1ContentBinding.inflate(inflater, container,false)
        val root: View = binding.root
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        val backBtn = binding.backBtn
        backBtn.setOnClickListener{
            (activity as MainActivity).navController.navigate(R.id.action_infoChapter1Content_to_infoAllChapters)
        }
    }

}