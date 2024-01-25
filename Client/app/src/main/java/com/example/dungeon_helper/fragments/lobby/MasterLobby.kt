package com.example.dungeon_helper.fragments.lobby

import android.content.Intent
import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.MasterLobbyActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentMasterLobbyBinding

class MasterLobby : Fragment() {

    companion object {
        fun newInstance() = MasterLobby()
    }

    private lateinit var viewModel: MasterLobbyViewModel
    private var _binding: FragmentMasterLobbyBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val masterLobbyViewModel = ViewModelProvider(this)[MasterLobbyViewModel::class.java]
        _binding = FragmentMasterLobbyBinding.inflate(inflater, container, false)
        val root: View = binding.root
        val textView: TextView = binding.textPlay
        masterLobbyViewModel.text.observe(viewLifecycleOwner){
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        val viewChar1 = binding.viewChar1
        viewChar1.setOnClickListener{
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobby_to_masterLobbyViewChar)
        }
        val viewChar2 = binding.viewChar2
        val viewChar3 = binding.viewChar3
        val infoBtn = binding.infoBtn
        infoBtn.setOnClickListener {
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobby_to_masterLobbyInfoAllChapters)
        }

        val actBtn = binding.actionBtn
        actBtn.setOnClickListener {
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobby_to_masterLobbyChoose)
        }

        val backBtn = binding.backBtn
        backBtn.setOnClickListener {
            val intent =
                Intent(activity as MasterLobbyActivity, MasterLobbyActivity::class.java)

            startActivity(intent)
        }

    }

}